package client

import (
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-podman/pkg/client/container"
)

func (c *Client) Container(enrich bool) []*container.Container {
	page, e := containers.List(
		c.context,
		&containers.ListOptions{All: ptr.To(true)},
	)
	errors.PanicOnError(e)

	result := container.NewSlice(page)

	if enrich {
		c.enrichMany(result)
	}

	return result
}
