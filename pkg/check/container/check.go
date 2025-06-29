package container

import (
	"fmt"
	"github.com/funtimecoding/go-podman/pkg/check/container/option"
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/client/container"
	"github.com/funtimecoding/go-podman/pkg/constant"
	"github.com/funtimecoding/go-podman/pkg/podman"
)

func Check(o *option.Container) {
	c := client.New()
	autoStart := podman.AutoStart()
	containers := c.Container(true)
	var relevant []*container.Container

	if o.All {
		relevant = containers
	} else {
		for _, n := range containers {
			if len(n.Concern) > 0 {
				relevant = append(relevant, n)
			}
		}
	}

	if o.Notation {
		printNotation(relevant, o)

		return
	}

	f := constant.Format

	for _, n := range relevant {
		fmt.Println(n.Format(f))

		if autoStart && !n.Running() && n.AlwaysRestart() {
			c.StartName(n.Name)
		}
	}
}
