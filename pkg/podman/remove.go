package podman

import (
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/domain/entities/reports"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/podman/container"
)

func (c *Client) Remove(o *container.Container) []*reports.RmReport {
	result, e := containers.Remove(
		c.context,
		o.Alias,
		&containers.RemoveOptions{},
	)
	errors.PanicOnError(e)

	return result
}
