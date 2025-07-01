package client

import (
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/client/image"
)

func (c *Client) Image() []*image.Image {
	result, e := images.List(c.context, &images.ListOptions{})
	errors.PanicOnError(e)

	return image.NewSlice(result)
}
