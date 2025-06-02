package client

import (
	"github.com/containers/podman/v5/pkg/api/handlers"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	containerType "github.com/docker/docker/api/types/container"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) session(
	container string,
	command ...string,
) string {
	result, e := containers.ExecCreate(
		c.context,
		container,
		&handlers.ExecCreateConfig{
			ExecOptions: containerType.ExecOptions{
				Cmd:          command,
				Tty:          false,
				AttachStdout: true,
				AttachStderr: true,
			},
		},
	)
	errors.PanicOnError(e)

	return result
}
