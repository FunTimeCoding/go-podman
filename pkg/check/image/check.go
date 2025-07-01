package image

import (
	"fmt"
	"github.com/funtimecoding/go-podman/pkg/check/image/option"
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/client/image"
	"github.com/funtimecoding/go-podman/pkg/constant"
	"github.com/funtimecoding/go-podman/pkg/podman"
)

func Check(o *option.Image) {
	// TODO: Checks all images for updates
	c := client.New()
	autoClean := podman.ImageAutoClean()
	f := constant.Format
	byName := make(map[string][]*image.Image)
	images := c.Image()

	for _, i := range images {
		byName[i.Name] = append(byName[i.Name], i)
	}

	for _, v := range byName {
		if len(v) > 1 {
			for _, i := range v {
				if autoClean && i.Dangling {
					c.RemoveImage([]string{i.Identifier})
				}
			}
		}
	}

	images = c.Image()

	if o.Notation {
		printNotation(images, o)

		return
	}

	for _, i := range images {
		fmt.Println(i.Format(f))
	}
}
