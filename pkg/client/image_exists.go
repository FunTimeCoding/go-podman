package client

import "github.com/funtimecoding/go-podman/pkg/container"

func (c *Client) ImageExists(o *container.Container) bool {
	for _, i := range c.Image(false) {
		if i.Name == o.Name {
			return true
		}
	}

	return false
}
