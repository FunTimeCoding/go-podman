package client

import (
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) StartName(name string) {
	errors.PanicOnError(
		containers.Start(c.context, name, &containers.StartOptions{}),
	)
}
