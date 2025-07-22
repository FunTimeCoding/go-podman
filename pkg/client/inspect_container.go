package client

import (
	"github.com/containers/podman/v5/libpod/define"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) InspectContainer(name string) *define.InspectContainerData {
	result, e := containers.Inspect(
		c.context,
		name,
		&containers.InspectOptions{},
	)
	errors.PanicOnError(e)

	return result
}
