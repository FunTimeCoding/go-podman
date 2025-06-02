package client

import (
	"github.com/funtimecoding/go-podman/pkg/container"
	"golang.org/x/exp/slices"
)

func (c *Client) ImageExists(o *container.Container) bool {
	for _, i := range c.Image() {
		if slices.Contains(i.Names, o.Name) {
			return true
		}
	}

	return false
}
