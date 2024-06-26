package compose

import (
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/container"
)

func New(c ...*container.Container) *Compose {
	return &Compose{podman: client.New(), containers: c}
}
