package client

import "github.com/funtimecoding/go-podman/pkg/client/volume"

func (c *Client) enrichManyVolume(v []*volume.Volume) []*volume.Volume {
	for _, o := range v {
		c.enrichOneVolume(o)
	}

	return v
}
