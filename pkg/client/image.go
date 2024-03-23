package client

import (
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/domain/entities"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Image() []*entities.ImageSummary {
	result, e := images.List(c.context, &images.ListOptions{})
	errors.PanicOnError(e)

	return result
}
