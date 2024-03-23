package podman

import (
	"github.com/funtimecoding/go-podman/pkg/podman/container"
	"golang.org/x/exp/slices"
)

func (c *Client) Running(o *container.Container) bool {
	var result bool

	for _, element := range c.Container() {
		if element.State != RunningState {
			continue
		}

		if slices.Contains(element.Names, o.Alias) {
			result = true

			break
		}
	}

	return result
}
