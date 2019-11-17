package spec_test

import (
	"testing"

	"code.cloudfoundry.org/garden"
	"github.com/concourse/concourse/worker/backend/spec"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	*require.Assertions
}

func (s *Suite) TestContainerSpecValidations() {
	for _, tc := range []struct {
		desc string
		spec garden.ContainerSpec
	}{
		{
			desc: "no handle specified",
			spec: garden.ContainerSpec{},
		},
		{
			desc: "rootfsPath not specified",
			spec: garden.ContainerSpec{
				Handle: "handle",
			},
		},
		{
			desc: "rootfsPath without scheme",
			spec: garden.ContainerSpec{
				Handle:     "handle",
				RootFSPath: "foo",
			},
		},
		{
			desc: "rootfsPath with unknown scheme",
			spec: garden.ContainerSpec{
				Handle:     "handle",
				RootFSPath: "weird://foo",
			},
		},
		{
			desc: "rootfsPath not being absolute",
			spec: garden.ContainerSpec{
				Handle:     "handle",
				RootFSPath: "raw://../not/absolute/at/all",
			},
		},
		{
			desc: "both rootfsPath and image specified",
			spec: garden.ContainerSpec{
				Handle:     "handle",
				RootFSPath: "foo",
				Image:      garden.ImageRef{URI: "bar"},
			},
		},
		{
			desc: "no rootfsPath, but image specified w/out scheme",
			spec: garden.ContainerSpec{
				Handle: "handle",
				Image:  garden.ImageRef{URI: "bar"},
			},
		},
		{
			desc: "no rootfsPath, but image specified w/ unknown scheme",
			spec: garden.ContainerSpec{
				Handle: "handle",
				Image:  garden.ImageRef{URI: "weird://bar"},
			},
		},
	} {
		s.T().Run(tc.desc, func(t *testing.T) {
			_, err := spec.OciSpec(tc.spec)
			s.Error(err)
		})
	}
}

func (s *Suite) TestOciSpecBindMounts() {
	for _, tc := range []struct {
		desc     string
		mounts   []garden.BindMount
		expected []specs.Mount
		succeeds bool
	}{
		{
			desc:     "unknown mode",
			succeeds: false,
			mounts: []garden.BindMount{
				{
					SrcPath: "/a",
					DstPath: "/b",
					Mode:    123,
					Origin:  garden.BindMountOriginHost,
				},
			},
		},
		{
			desc:     "unknown origin",
			succeeds: false,
			mounts: []garden.BindMount{
				{
					SrcPath: "/a",
					DstPath: "/b",
					Mode:    garden.BindMountModeRO,
					Origin:  123,
				},
			},
		},
		{
			desc:     "w/out src",
			succeeds: false,
			mounts: []garden.BindMount{
				{
					DstPath: "/b",
					Mode:    garden.BindMountModeRO,
					Origin:  garden.BindMountOriginHost,
				},
			},
		},
		{
			desc:     "non-absolute src",
			succeeds: false,
			mounts: []garden.BindMount{
				{
					DstPath: "/b",
					Mode:    garden.BindMountModeRO,
					Origin:  garden.BindMountOriginHost,
				},
			},
		},
		{
			desc:     "w/out dest",
			succeeds: false,
			mounts: []garden.BindMount{
				{
					SrcPath: "/a",
					Mode:    garden.BindMountModeRO,
					Origin:  garden.BindMountOriginHost,
				},
			},
		},
		{
			desc:     "non-absolute dest",
			succeeds: false,
			mounts: []garden.BindMount{
				{
					DstPath: "/b",
					Mode:    garden.BindMountModeRO,
					Origin:  garden.BindMountOriginHost,
				},
			},
		},
	} {
		s.T().Run(tc.desc, func(t *testing.T) {
			actual, err := spec.OciSpecBindMounts(tc.mounts)
			if !tc.succeeds {
				s.Error(err)
				return
			}

			s.NoError(err)
			s.Equal(tc.expected, actual)
		})
	}
}

// func (s *Suite) TestLinuxContainerizationSpec() {
// 	for _, tc := range []struct {
// 		desc     string
// 		gdn      garden.ContainerSpec
// 		expected *specs.Spec
// 	}{
// 		{
// 			desc: "privileged",
// 			gdn: garden.ContainerSpec{
// 				Handle:     "handle",
// 				RootFSPath: "raw:///rootfs",
// 			},
// 			expected: &specs.Spec{
// 				Hostname: "handle",
// 				Linux: &specs.Linux{
// 					Namespaces: spec.PrivilegedContainerNamespaces,
// 					Resources:  &specs.LinuxResources{Devices: spec.AnyContainerDevices},
// 				},
// 				Process: &specs.Process{
// 					Capabilities: &spec.PrivilegedContainerCapabilities,
// 				},
// 				Root:    &specs.Root{Path: "/rootfs"},
// 				Version: specs.Version,
// 			},
// 		},
// 	} {
// 		s.T().Run(tc.desc, func(t *testing.T) {
// 			actual, err := spec.OciSpec(tc.gdn)
// 			s.NoError(err)
// 			s.Equal(tc.expected, actual)
// 		})
// 	}
// }

func (s *Suite) TestOciNamespaces() {
	for _, tc := range []struct {
		desc       string
		privileged bool
		expected   []specs.LinuxNamespace
	}{
		{
			desc:       "privileged",
			privileged: true,
			expected:   spec.PrivilegedContainerNamespaces,
		},
		{
			desc:       "unprivileged",
			privileged: false,
			expected:   spec.UnprivilegedContainerNamespaces,
		},
	} {
		s.T().Run(tc.desc, func(t *testing.T) {
			s.Equal(tc.expected, spec.OciNamespaces(tc.privileged))
		})
	}
}

func (s *Suite) TestOciCapabilities() {
	for _, tc := range []struct {
		desc       string
		privileged bool
		expected   specs.LinuxCapabilities
	}{
		{
			desc:       "privileged",
			privileged: true,
			expected:   spec.PrivilegedContainerCapabilities,
		},
		{
			desc:       "unprivileged",
			privileged: false,
			expected:   spec.UnprivilegedContainerCapabilities,
		},
	} {
		s.T().Run(tc.desc, func(t *testing.T) {
			s.Equal(tc.expected, spec.OciCapabilities(tc.privileged))
		})
	}
}

func (s *Suite) TestContainerSpec() {
	for _, tc := range []struct {
		desc  string
		gdn   garden.ContainerSpec
		check func(*specs.Spec)
	}{

		{
			desc: "env",
			gdn: garden.ContainerSpec{
				Handle: "handle", RootFSPath: "raw:///rootfs",
				Env: []string{"foo=bar"},
			},
			check: func(oci *specs.Spec) {
				s.Equal([]string{"foo=bar"}, oci.Process.Env)
			},
		},
		{
			desc: "mounts",
			gdn: garden.ContainerSpec{
				Handle: "handle", RootFSPath: "raw:///rootfs",
				BindMounts: []garden.BindMount{
					{ // ro mount
						SrcPath: "/a",
						DstPath: "/b",
						Mode:    garden.BindMountModeRO,
						Origin:  garden.BindMountOriginHost,
					},
					{ // rw mount
						SrcPath: "/a",
						DstPath: "/b",
						Mode:    garden.BindMountModeRW,
						Origin:  garden.BindMountOriginHost,
					},
				},
			},
			check: func(oci *specs.Spec) {
				s.Equal([]specs.Mount{
					{
						Source:      "/a",
						Destination: "/b",
						Type:        "bind",
						Options:     []string{"bind", "ro"},
					},
					{
						Source:      "/a",
						Destination: "/b",
						Type:        "bind",
						Options:     []string{"bind", "rw"},
					},
				}, oci.Mounts)
			},
		},
		// {
		// desc: "limits",
		// gdn:  garden.ContainerSpec{},
		// check: func(oci *specs.Spec) {
		// 	s.Equal([]string{"foo=bar"}, oci.Process.Env)
		// },
		// },
		// {
		// 	desc: "properties",
		// 	gdn:  garden.ContainerSpec{},
		// 	check: func(oci *specs.Spec) {
		// 		s.Equal([]string{"foo=bar"}, oci.Process.Env)
		// 	},
		// },
	} {
		s.T().Run(tc.desc, func(t *testing.T) {
			actual, err := spec.OciSpec(tc.gdn)
			s.NoError(err)

			tc.check(actual)
		})
	}
}

func TestSuite(t *testing.T) {
	suite.Run(t, &Suite{
		Assertions: require.New(t),
	})
}
