package podman

import (
	"github.com/containers/common/libnetwork/types"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/containers/podman/v4/pkg/specgen"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/podman/container"
)

func (c *Client) Create(
	o *container.Container,
) entities.ContainerCreateResponse {
	g := specgen.NewSpecGenerator(o.Name, false)
	g.Name = o.Alias
	g.Remove = true

	if o.Port > 0 {
		port := uint16(o.Port)
		g.PublishExposedPorts = true
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
