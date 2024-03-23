package compose

import (
	"github.com/funtimecoding/go-podman/pkg/podman"
	"github.com/funtimecoding/go-podman/pkg/podman/container"
)

func New(c ...*container.Container) *Tester {
	return &Tester{podman: podman.New(), containers: c}
}
