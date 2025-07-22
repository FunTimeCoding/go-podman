package client

import "github.com/funtimecoding/go-podman/pkg/client/container"

func (c *Client) enrichOneContainer(o *container.Container) {
	o.Enrich(c.InspectContainer(o.Name))
}
