package compose

import "fmt"

func (c *Compose) Create() {
	for _, element := range c.containers {
		if !c.podman.ImageExists(element) {
			c.podman.Pull(element)
		}

		if !c.podman.Exists(element) {
			r := c.podman.Create(element)

			if len(r.Warnings) > 0 {
				fmt.Printf("Identifier: %s\n", r.ID)

				for _, warning := range r.Warnings {
					fmt.Printf("Warning: %s\n", warning)
				}
			}
		}

		if !c.podman.Running(element) {
			c.podman.Start(element)
		}
	}
}
