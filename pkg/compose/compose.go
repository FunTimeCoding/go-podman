package compose

import (
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/container"
)

type Compose struct {
	podman     *client.Client
	containers []*container.Container
}
