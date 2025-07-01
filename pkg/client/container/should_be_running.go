package container

func (c *Container) ShouldBeRunning() bool {
	return !c.Running() && c.AlwaysRestart()
}
