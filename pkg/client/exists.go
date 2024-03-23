package client

import (
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/container"
)

func (c *Client) Exists(o *container.Container) bool {
	result, e := containers.Exists(
		c.context,
		o.Alias,
		&containers.ExistsOptions{},
	)
	errors.PanicOnError(e)

	return result
}
