package compose

func (c *Compose) Create() {
	for _, element := range c.containers {
		if !c.podman.ImageExists(element) {
			c.podman.Pull(element)
		}

		if !c.podman.Exists(element) {
			c.podman.Create(element)
		}

		if !c.podman.Running(element) {
			c.podman.Start(element)
		}
	}
}
