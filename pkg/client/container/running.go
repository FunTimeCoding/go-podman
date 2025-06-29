package container

func (c *Container) Running() bool {
	return c.State == RunningState
}
