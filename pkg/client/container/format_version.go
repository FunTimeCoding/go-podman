package container

func (c *Container) formatVersion() string {
	result := c.Version

	if len(result) > 12 {
		result = result[:12]
	}

	return result
}
