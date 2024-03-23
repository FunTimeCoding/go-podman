package podman

import (
	"github.com/containers/podman/v4/pkg/bindings/images"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/podman/container"
)

func (c *Client) Pull(o *container.Container) []string {
	result, e := images.Pull(c.context, o.Name, &images.PullOptions{})
	errors.PanicOnError(e)

	return result
}
