package example

import (
	"fmt"
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func Container() {
	c := client.New()
	f := constant.Format.Raw()

	for _, o := range c.Container(false) {
		fmt.Printf("Container: %s\n", o.Format(f))
	}
}
