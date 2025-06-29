package container

func (c *Container) AlwaysRestart() bool {
	return c.Restart == Always
}
