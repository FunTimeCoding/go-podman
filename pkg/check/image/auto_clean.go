package image

import (
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/client/image"
	"github.com/funtimecoding/go-podman/pkg/podman"
)

func autoClean(c *client.Client) {
	if !podman.ImageAutoClean() {
		return
	}

	byName := make(map[string][]*image.Image)

	for _, i := range c.Image(true) {
		byName[i.Name] = append(byName[i.Name], i)
	}

	for _, v := range byName {
		if len(v) < 2 {
			continue
		}

		for _, i := range v {
			if i.ShouldDelete() {
				c.RemoveImage([]string{i.Identifier})
			}
		}
	}
}
