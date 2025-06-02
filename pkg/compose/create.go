package compose

import "fmt"

func (c *Compose) Create() {
	for _, o := range c.containers {
		if !c.podman.ImageExists(o) {
			c.podman.Pull(o)
		}

		if !c.podman.Exists(o) {
			r := c.podman.Create(o)

			if len(r.Warnings) > 0 {
				fmt.Printf("Identifier: %s\n", r.ID)

				for _, warning := range r.Warnings {
					fmt.Printf("Warning: %s\n", warning)
				}
			}
		}

		if !c.podman.Running(o) {
			c.podman.Start(o)
		}
	}
}
