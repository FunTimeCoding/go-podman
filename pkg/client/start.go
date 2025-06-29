package client

import "github.com/funtimecoding/go-podman/pkg/container"

func (c *Client) Start(o *container.Container) {
	c.StartName(o.Alias)
}
