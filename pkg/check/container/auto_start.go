package container

import (
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/podman"
)

func autoStart(c *client.Client) {
	if !podman.ContainerAutoStart() {
		return
	}

	for _, n := range c.Container(true) {
		if n.ShouldBeRunning() {
			c.StartName(n.Name)
		}
	}
}
