package client

import (
	"github.com/funtimecoding/go-podman/pkg/container"
	"golang.org/x/exp/slices"
)

func (c *Client) Running(o *container.Container) bool {
	var result bool

	for _, n := range c.Container() {
		if n.State != RunningState {
			continue
		}

		if slices.Contains(n.Names, o.Alias) {
			result = true

			break
		}
	}

	return result
}
