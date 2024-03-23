package compose

func (c *Compose) Destroy() {
	for _, element := range c.containers {
		if !c.podman.Exists(element) {
			return
		}

		if c.podman.Running(element) {
			c.podman.Stop(element)
		}

		c.podman.Remove(element)
	}
}
