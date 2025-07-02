package container

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-podman/pkg/check/container/option"
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func Check(o *option.Container) {
	c := client.New()
	autoStart(c)
	containers := monitor.OnlyConcerns(c.Container(true), o.All)

	if o.Notation {
		printNotation(containers, o)

		return
	}

	f := constant.Format

	for _, n := range containers {
		fmt.Println(n.Format(f))
	}
}
