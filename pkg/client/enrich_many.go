package client

import "github.com/funtimecoding/go-podman/pkg/client/container"

func (c *Client) enrichMany(v []*container.Container) []*container.Container {
	for _, o := range v {
		c.enrichOne(o)
		o.Validate()
	}

	return v
}
