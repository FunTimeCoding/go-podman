package client

import (
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/container"
)

func (c *Client) Stop(o *container.Container) {
	errors.PanicOnError(
		containers.Stop(c.context, o.Alias, &containers.StopOptions{}),
	)
}
