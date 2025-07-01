package client

import "github.com/funtimecoding/go-podman/pkg/client/image"

func (c *Client) enrichManyImage(v []*image.Image) []*image.Image {
	for _, o := range v {
		c.enrichOneImage(o)
	}

	return v
}
