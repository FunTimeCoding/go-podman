package container

import "slices"

func (c *Container) HasConcern(s string) bool {
	return slices.Contains(c.concern, s)
}
