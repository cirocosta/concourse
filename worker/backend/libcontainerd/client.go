package libcontainerd

import (
	"context"
	"fmt"

	"github.com/containerd/containerd"
	"github.com/containerd/containerd/namespaces"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . Client

// Client represents the minimum interface used to communicate with containerd
// to manage containers.
//
type Client interface {

	// Init provides the initialization of internal structures necessary by
	// the client, e.g., instantiation of the gRPC client.
	//
	Init() (err error)

	// Version queries containerd's version service in order to verify
	// connectivity.
	//
	Version(ctx context.Context) (err error)

	// Stop deallocates any initialization performed by `Init()` and
	// subsequent calls to methods of this interface.
	//
	Stop() (err error)

	// NewContainer creates a container in containerd.
	//
	NewContainer(
		ctx context.Context, id string, opts ...containerd.NewContainerOpts,
	) (
		container containerd.Container, err error,
	)
}

type client struct {
	addr      string
	namespace string

	containerd *containerd.Client
}

func New(addr, namespace string) *client {
	return &client{
		addr:      addr,
		namespace: namespace,
	}
}

func (c *client) Init() (err error) {
	c.containerd, err = containerd.New(c.addr)
	if err != nil {
		err = fmt.Errorf("failed to connect to addr %s: %w", c.addr, err)
		return
	}

	return
}

func (c *client) Stop() (err error) {
	if c.containerd == nil {
		return
	}

	err = c.containerd.Close()
	return
}

func (c *client) NewContainer(
	ctx context.Context, id string, opts ...containerd.NewContainerOpts,
) (
	containerd.Container, error,
) {
	ctx = namespaces.WithNamespace(ctx, c.namespace)

	// create a snapshot of that rootfs

	return c.containerd.NewContainer(ctx, id, opts...)
}

func (c *client) Version(ctx context.Context) (err error) {
	_, err = c.containerd.Version(ctx)
	return
}
