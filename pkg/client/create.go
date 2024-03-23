package client

import (
	"github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/domain/entities"
	"github.com/containers/podman/v5/pkg/machine/ignition"
	"github.com/containers/podman/v5/pkg/specgen"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/container"
)

func (c *Client) Create(
	o *container.Container,
) entities.ContainerCreateResponse {
	g := specgen.NewSpecGenerator(o.Name, false)
	g.Name = o.Alias
	g.Remove = ignition.BoolToPtr(true)

	if o.Port > 0 {
		port := uint16(o.Port)
		g.PublishExposedPorts = ignition.BoolToPtr(true)
		g.PortMappings = append(
			g.PortMappings,
			types.PortMapping{ContainerPort: port, HostPort: port},
		)
		g.NetNS = specgen.Namespace{NSMode: specgen.Bridge}
		g.Expose = make(map[uint16]string)
		g.Expose[port] = TCPPort
	}

	if len(o.Environment) > 0 {
		g.Env = o.Environment
	}

	result, e := containers.CreateWithSpec(
		c.context,
		g,
		&containers.CreateOptions{},
	)
	errors.PanicOnError(e)

	return result
}
