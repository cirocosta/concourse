package spec

import (
	"fmt"
	"path/filepath"
	"strings"

	"code.cloudfoundry.org/garden"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

var (

	// Most of these defaults have been taken from Guardian's code base.
	//
	// ref: https://github.com/cloudfoundry/guardian/blob/6b021168907b2f2ae25cf54acae8b454b430332f/guardiancmd/server.go
	// ref: https://github.com/containerd/containerd/blob/ec661e8ceb85bca68cd759f9d9513cb6f103ca42/oci/spec.go#L112
	// ref: https://github.com/cloudfoundry/guardian/blob/0a658a3e51595c214b0e0cb43b4133d274011b44/guardiancmd/command.go
	//

	AnyContainerDevices = []specs.LinuxDeviceCgroup{
		// runc allows these
		{Access: "m", Type: "c", Major: deviceWildcard(), Minor: deviceWildcard(), Allow: true},
		{Access: "m", Type: "b", Major: deviceWildcard(), Minor: deviceWildcard(), Allow: true},

		{Access: "rwm", Type: "c", Major: intRef(1), Minor: intRef(3), Allow: true},          // /dev/null
		{Access: "rwm", Type: "c", Major: intRef(1), Minor: intRef(8), Allow: true},          // /dev/random
		{Access: "rwm", Type: "c", Major: intRef(1), Minor: intRef(7), Allow: true},          // /dev/full
		{Access: "rwm", Type: "c", Major: intRef(5), Minor: intRef(0), Allow: true},          // /dev/tty
		{Access: "rwm", Type: "c", Major: intRef(1), Minor: intRef(5), Allow: true},          // /dev/zero
		{Access: "rwm", Type: "c", Major: intRef(1), Minor: intRef(9), Allow: true},          // /dev/urandom
		{Access: "rwm", Type: "c", Major: intRef(5), Minor: intRef(1), Allow: true},          // /dev/console
		{Access: "rwm", Type: "c", Major: intRef(136), Minor: deviceWildcard(), Allow: true}, // /dev/pts/*
		{Access: "rwm", Type: "c", Major: intRef(5), Minor: intRef(2), Allow: true},          // /dev/ptmx
		{Access: "rwm", Type: "c", Major: intRef(10), Minor: intRef(200), Allow: true},       // /dev/net/tun

		// we allow these
		{Access: "rwm", Type: "c", Major: intRef(10), Minor: intRef(229), Allow: true}, // /dev/fuse
	}

	PrivilegedContainerCapabilities = specs.LinuxCapabilities{
		Effective:   privilegedCaps,
		Bounding:    privilegedCaps,
		Inheritable: privilegedCaps,
		Permitted:   privilegedCaps,
	}

	UnprivilegedContainerCapabilities = specs.LinuxCapabilities{
		Effective:   unprivilegedCaps,
		Bounding:    unprivilegedCaps,
		Inheritable: unprivilegedCaps,
		Permitted:   unprivilegedCaps,
	}

	PrivilegedContainerNamespaces = []specs.LinuxNamespace{
		{Type: specs.PIDNamespace},
		{Type: specs.IPCNamespace},
		{Type: specs.UTSNamespace},
		{Type: specs.MountNamespace},
		{Type: specs.NetworkNamespace},
	}

	UnprivilegedContainerNamespaces = append(PrivilegedContainerNamespaces,
		specs.LinuxNamespace{Type: specs.UserNamespace},
	)

	unprivilegedCaps = []string{
		"CAP_CHOWN",
		"CAP_DAC_OVERRIDE",
		"CAP_FSETID",
		"CAP_FOWNER",
		"CAP_MKNOD",
		"CAP_NET_RAW",
		"CAP_SETGID",
		"CAP_SETUID",
		"CAP_SETFCAP",
		"CAP_SETPCAP",
		"CAP_NET_BIND_SERVICE",
		"CAP_SYS_CHROOT",
		"CAP_KILL",
		"CAP_AUDIT_WRITE",
	}

	privilegedCaps = []string{
		"CAP_AUDIT_CONTROL",
		"CAP_AUDIT_READ",
		"CAP_AUDIT_WRITE",
		"CAP_BLOCK_SUSPEND",
		"CAP_CHOWN",
		"CAP_DAC_OVERRIDE",
		"CAP_DAC_READ_SEARCH",
		"CAP_FOWNER",
		"CAP_FSETID",
		"CAP_IPC_LOCK",
		"CAP_IPC_OWNER",
		"CAP_KILL",
		"CAP_LEASE",
		"CAP_LINUX_IMMUTABLE",
		"CAP_MAC_ADMIN",
		"CAP_MAC_OVERRIDE",
		"CAP_MKNOD",
		"CAP_NET_ADMIN",
		"CAP_NET_BIND_SERVICE",
		"CAP_NET_BROADCAST",
		"CAP_NET_RAW",
		"CAP_SETGID",
		"CAP_SETFCAP",
		"CAP_SETPCAP",
		"CAP_SETUID",
		"CAP_SYS_ADMIN",
		"CAP_SYS_BOOT",
		"CAP_SYS_CHROOT",
		"CAP_SYS_MODULE",
		"CAP_SYS_NICE",
		"CAP_SYS_PACCT",
		"CAP_SYS_PTRACE",
		"CAP_SYS_RAWIO",
		"CAP_SYS_RESOURCE",
		"CAP_SYS_TIME",
		"CAP_SYS_TTY_CONFIG",
		"CAP_SYSLOG",
		"CAP_WAKE_ALARM",
	}
)

func intRef(i int64) *int64  { return &i }
func deviceWildcard() *int64 { return intRef(-1) }

// OciSpec converts a given `garden` container specification to an OCI spec.
//
func OciSpec(gdn garden.ContainerSpec) (oci *specs.Spec, err error) {
	var (
		rootfs string
		mounts []specs.Mount

		namespaces   = OciNamespaces(gdn.Privileged)
		capabilities = OciCapabilities(gdn.Privileged)
	)

	if gdn.RootFSPath == "" {
		gdn.RootFSPath = gdn.Image.URI
	}

	rootfs, err = rootfsDir(gdn.RootFSPath)
	if err != nil {
		return
	}

	mounts, err = OciSpecBindMounts(gdn.BindMounts)
	if err != nil {
		return
	}

	oci = &specs.Spec{
		Version:  specs.Version,
		Hostname: gdn.Handle,
		Process: &specs.Process{
			Capabilities: &capabilities,
			Env:          gdn.Env,
		},
		Root:        &specs.Root{Path: rootfs},
		Mounts:      mounts,
		Annotations: map[string]string(gdn.Properties),
		Linux: &specs.Linux{
			Namespaces: namespaces,
			Resources: &specs.LinuxResources{
				Devices: AnyContainerDevices,
				// Memory:  nil,
				// Cpu:     nil,
			},
		},
	}

	// deals with
	// - limits
	// - masked paths
	// - rootfs propagation
	// - seccomp
	// - user namespaces: uid/gid mappings
	// x capabilities
	// x devices
	// x env
	// x hostname
	// x mounts
	// x namespaces
	// x rootfs

	return
}

func OciNamespaces(privileged bool) []specs.LinuxNamespace {
	if !privileged {
		return UnprivilegedContainerNamespaces
	}

	return PrivilegedContainerNamespaces
}

func OciCapabilities(privileged bool) specs.LinuxCapabilities {
	if !privileged {
		return UnprivilegedContainerCapabilities
	}

	return PrivilegedContainerCapabilities
}

// OciSpecBindMounts converts garden bindmounts to oci spec mounts.
//
func OciSpecBindMounts(bindMounts []garden.BindMount) (mounts []specs.Mount, err error) {
	for _, bindMount := range bindMounts {
		if bindMount.SrcPath == "" || bindMount.DstPath == "" {
			err = fmt.Errorf("src and dst must not be empty")
			return
		}

		if !filepath.IsAbs(bindMount.SrcPath) || !filepath.IsAbs(bindMount.DstPath) {
			err = fmt.Errorf("src and dst must be absolute")
			return
		}

		if bindMount.Origin != garden.BindMountOriginHost {
			err = fmt.Errorf("unknown bind mount origin %d", bindMount.Origin)
			return
		}

		mode := "ro"
		switch bindMount.Mode {
		case garden.BindMountModeRO:
		case garden.BindMountModeRW:
			mode = "rw"
		default:
			err = fmt.Errorf("unknown bind mount mode %d", bindMount.Mode)
			return
		}

		mounts = append(mounts, specs.Mount{
			Source:      bindMount.SrcPath,
			Destination: bindMount.DstPath,
			Type:        "bind",
			Options:     []string{"bind", mode},
		})
	}

	return
}

// rootfsDir takes a raw rootfs uri and extracts the directory that it points to,
// if using a valid scheme (`raw://`)
//
func rootfsDir(raw string) (directory string, err error) {
	if raw == "" {
		err = fmt.Errorf("rootfs must not be empty")
		return
	}

	parts := strings.SplitN(raw, "://", 2)
	if len(parts) != 2 {
		err = fmt.Errorf("malformatted rootfs: must be of form 'scheme://<abs_dir>'")
		return
	}

	var scheme string
	scheme, directory = parts[0], parts[1]
	if scheme != "raw" {
		err = fmt.Errorf("unsupported scheme '%s'", scheme)
		return
	}

	if !filepath.IsAbs(directory) {
		err = fmt.Errorf("directory must be an absolute path")
		return
	}

	return
}
