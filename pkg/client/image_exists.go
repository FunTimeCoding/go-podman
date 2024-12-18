package client

import (
	"github.com/funtimecoding/go-podman/pkg/container"
	"golang.org/x/exp/slices"
)

func (c *Client) ImageExists(o *container.Container) bool {
	for _, element := range c.Image() {
		if slices.Contains(element.Names, o.Name) {
			return true
		}
	}

	return false
}
