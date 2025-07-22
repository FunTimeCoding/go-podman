package volume

import (
	"fmt"
	"github.com/funtimecoding/go-podman/pkg/check/volume/option"
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func Check(o *option.Volume) {
	c := client.New()
	elements := c.Volume()

	if o.Notation {
		return
	}

	f := constant.Format
	f.Raw()

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}
}
