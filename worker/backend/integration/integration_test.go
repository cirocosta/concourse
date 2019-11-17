package integration_test

import (
	"path/filepath"
	"testing"

	"code.cloudfoundry.org/garden"
	"github.com/concourse/concourse/worker/backend"
	"github.com/concourse/concourse/worker/backend/libcontainerd"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type BackendSuite struct {
	suite.Suite
	*require.Assertions

	backend backend.Backend
	client  *libcontainerd.Client
}

func (s *BackendSuite) SetupTest() {
	s.backend = backend.New(
		libcontainerd.New("/run/containerd/containerd.sock"),
	)

	s.NoError(s.backend.Start())
}

func (s *BackendSuite) TearDownTest() {
	s.backend.Stop()
}

func (s *BackendSuite) TestPing() {
	s.NoError(s.backend.Ping())
}

func (s *BackendSuite) TestContainerCreation() {
	handle := mustCreateHandle()
	rootfs, err := filepath.Abs("testdata/rootfs")
	s.NoError(err)

	_, err = s.backend.Create(garden.ContainerSpec{
		Handle:     handle,
		RootFSPath: "raw://" + rootfs,
	})
	s.NoError(err)

	defer s.backend.Destroy(handle)

	containers, err := s.backend.Containers(nil)
	s.NoError(err)

	s.Len(containers, 1)
}

func TestSuite(t *testing.T) {
	suite.Run(t, &BackendSuite{
		Assertions: require.New(t),
	})
}

func mustCreateHandle() string {
	u4, err := uuid.NewV4()
	if err != nil {
		panic("couldn't create new uuid")
	}

	return u4.String()
}
