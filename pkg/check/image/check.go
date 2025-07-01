package image

import (
	"fmt"
	"github.com/funtimecoding/go-podman/pkg/check/image/option"
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func Check(o *option.Image) {
	c := client.New()
	autoUpdate(c)
	autoClean(c)
	images := c.Image(true)

	if o.Notation {
		printNotation(images, o)

		return
	}

	f := constant.Format

	for _, i := range images {
		fmt.Println(i.Format(f))
	}
}
