package client

import "github.com/funtimecoding/go-podman/pkg/client/container"

func (c *Client) enrichOne(o *container.Container) {
	detail := c.Inspect(o.Name)
	o.Enrich(detail)
}
