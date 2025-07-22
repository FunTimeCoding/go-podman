package client

import (
	"github.com/containers/podman/v5/pkg/bindings/volumes"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/client/volume"
)

func (c *Client) Volume() []*volume.Volume {
	page, e := volumes.List(c.context, &volumes.ListOptions{})
	errors.PanicOnError(e)
	result := volume.NewSlice(page)

	return result
}
