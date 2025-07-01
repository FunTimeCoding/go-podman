package example

import (
	"fmt"
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func Image() {
	c := client.New()
	f := constant.Format

	for _, i := range c.Image() {
		fmt.Printf("Image: %s\n", i.Format(f))
	}
}
