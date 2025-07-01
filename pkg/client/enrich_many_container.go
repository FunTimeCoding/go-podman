package client

import "github.com/funtimecoding/go-podman/pkg/client/container"

func (c *Client) enrichManyContainer(v []*container.Container) []*container.Container {
	for _, o := range v {
		c.enrichOneContainer(o)
	}

	return v
}
