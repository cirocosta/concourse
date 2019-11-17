package spec

import "github.com/opencontainers/runtime-spec/specs-go"

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
)

func intRef(i int64) *int64  { return &i }
func deviceWildcard() *int64 { return intRef(-1) }
