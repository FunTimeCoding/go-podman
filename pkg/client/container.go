package client

import (
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/domain/entities"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/ptr"
	"github.com/funtimecoding/go-podman/pkg/client/container"
)

func (c *Client) Container() []entities.ListContainer {
	result, e := containers.List(c.context, &containers.ListOptions{})
	errors.PanicOnError(e)

	return result
}

func (c *Client) Container2() []*container.Container {
	result, e := containers.List(
		c.context,
		&containers.ListOptions{All: ptr.To(true)},
	)
	errors.PanicOnError(e)

	return container.NewSlice(result)
}
