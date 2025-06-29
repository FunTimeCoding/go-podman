package client

import (
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/domain/entities"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/client/image"
)

func (c *Client) Image() []*entities.ImageSummary {
	result, e := images.List(c.context, &images.ListOptions{})
	errors.PanicOnError(e)

	return result
}

func (c *Client) Image2() []*image.Image {
	result, e := images.List(c.context, &images.ListOptions{})
	errors.PanicOnError(e)

	return image.NewSlice(result)
}
