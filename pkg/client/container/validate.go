package container

func (c *Container) Validate() {
	if c.AlwaysRestart() && !c.Running() {
		c.Concern = append(c.Concern, "should_run")
	}
}
