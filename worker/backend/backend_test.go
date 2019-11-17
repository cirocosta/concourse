package backend_test

import (
	"errors"
	"path/filepath"
	"testing"

	"code.cloudfoundry.org/garden"
	"github.com/concourse/concourse/worker/backend"
	"github.com/concourse/concourse/worker/backend/libcontainerd/libcontainerdfakes"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BackendSuite struct {
	suite.Suite
	*require.Assertions

	backend backend.Backend
	client  *libcontainerdfakes.FakeClient
}

func (s *BackendSuite) SetupTest() {
	s.client = new(libcontainerdfakes.FakeClient)
	s.backend = backend.New(s.client)
}

func (s *BackendSuite) TestPing() {
	for _, tc := range []struct {
		desc          string
		versionReturn error
		succeeds      bool
	}{
		{
			desc:          "fail from containerd version service",
			succeeds:      true,
			versionReturn: nil,
		},
		{
			desc:          "ok from containerd's version service",
			succeeds:      false,
			versionReturn: errors.New("errr"),
		},
	} {
		s.T().Run(tc.desc, func(t *testing.T) {
			s.client.VersionReturns(tc.versionReturn)

			err := s.backend.Ping()
			if tc.succeeds {
				s.NoError(err)
				return
			}

			s.Error(err)
		})
	}
}

func (s *BackendSuite) TestCreateBehavior() {
	// [cc] verify that it validates?

	rootfs, err := filepath.Abs("testdata/rootfs")
	s.NoError(err)

	spec := garden.ContainerSpec{
		Handle:     "handle",
		RootFSPath: "raw://" + rootfs,
	}

	_, err = s.backend.Create(spec)
	s.NoError(err)

	s.Equal(1, s.client.NewContainerCallCount())
}

func (s *BackendSuite) TestStart() {
	s.backend.Start()
	s.Equal(1, s.client.InitCallCount())
}

func (s *BackendSuite) TestStop() {
	s.backend.Stop()
	s.Equal(1, s.client.StopCallCount())
}

func TestSuite(t *testing.T) {
	suite.Run(t, &BackendSuite{
		Assertions: require.New(t),
	})
}
