package compose

import (
	"github.com/funtimecoding/go-podman/pkg/podman"
	"github.com/funtimecoding/go-podman/pkg/podman/container"
)

type Tester struct {
	podman     *podman.Client
	containers []*container.Container
}
