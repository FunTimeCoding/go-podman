package client

import (
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Container() []entities.ListContainer {
	result, e := containers.List(c.context, &containers.ListOptions{})
	errors.PanicOnError(e)

	return result
}
