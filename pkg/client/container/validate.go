package container

import "slices"

func (c *Container) Validate() {
	if c.ShouldBeRunning() && !slices.Contains(c.concern, ShouldRun) {
		c.concern = append(c.concern, ShouldRun)
	}
}
