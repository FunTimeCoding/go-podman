package client

import (
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/container"
)

func (c *Client) Start(o *container.Container) {
	errors.PanicOnError(
		containers.Start(c.context, o.Alias, &containers.StartOptions{}),
	)
}
