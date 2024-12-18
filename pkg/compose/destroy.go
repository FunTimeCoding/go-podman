package compose

func (c *Compose) Destroy() {
	for _, element := range c.containers {
		if !c.podman.Exists(element) {
			continue
		}

		if c.podman.Running(element) {
			c.podman.Stop(element)
		}

		if c.podman.Exists(element) {
			// If the remove flag was not set, it will be removed now
			c.podman.Remove(element)
		}
	}
}
