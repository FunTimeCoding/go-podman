package client

import "github.com/funtimecoding/go-podman/pkg/client/image"

func (c *Client) enrichOneImage(o *image.Image) {
	o.Enrich()
}
