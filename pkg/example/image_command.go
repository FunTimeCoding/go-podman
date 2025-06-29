package example

import (
	"fmt"
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/time"
	"github.com/funtimecoding/go-podman/pkg/command"
)

func ImageCommand() {
	for _, i := range command.Images() {
		name, version := i.NameVersion()
		fmt.Printf("Name: %s\n", name)
		fmt.Printf("Version: %s\n", version)
		fmt.Printf(
			"Create: %s\n",
			i.CreatedAt.Format(time.DateMinute),
		)
		fmt.Printf("Size: %s\n", units.HumanSize(float64(i.Size)))
	}
}
