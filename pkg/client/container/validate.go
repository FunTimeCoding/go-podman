package container

func (c *Container) Validate() {
	if c.ShouldBeRunning() {
		c.concern = append(c.concern, ShouldRun)
	}
}
