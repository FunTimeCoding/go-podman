package client

import "github.com/funtimecoding/go-podman/pkg/container"

func (c *Client) Running(o *container.Container) bool {
	var result bool

	for _, n := range c.Container(false) {
		if !n.Running() {
			continue
		}

		if n.Name == o.Alias {
			result = true

			break
		}
	}

	return result
}
