package client

import (
	"github.com/containers/podman/v5/pkg/bindings/volumes"
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) InspectVolume(name string) *types.VolumeConfigResponse {
	result, e := volumes.Inspect(
		c.context,
		name,
		&volumes.InspectOptions{},
	)
	errors.PanicOnError(e)

	return result
}
