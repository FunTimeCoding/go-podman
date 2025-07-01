package compose

func (c *Compose) Destroy() {
	for _, o := range c.containers {
		if !c.podman.Exists(o) {
			continue
		}

		if c.podman.Running(o) {
			c.podman.Stop(o)
		}

		if c.podman.Exists(o) {
			// If the remove flag was not set, it will be removed now
			c.podman.Remove(o.Alias)
		}
	}
}
