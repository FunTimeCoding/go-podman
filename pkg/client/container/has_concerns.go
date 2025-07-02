package container

func (c *Container) HasConcerns() bool {
	return len(c.concern) > 0
}
