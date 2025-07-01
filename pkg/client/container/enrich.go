package container

import "github.com/containers/podman/v5/libpod/define"

func (c *Container) Enrich(v *define.InspectContainerData) {
	c.RawDetail = v
	c.Restart = v.HostConfig.RestartPolicy.Name
	c.Validate()
}
