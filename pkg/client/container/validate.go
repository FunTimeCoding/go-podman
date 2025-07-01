package container

func (c *Container) Validate() {
	if c.ShouldBeRunning() {
		c.Concern = append(c.Concern, ShouldRun)
	}
}
