package container

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"time"
)

type Container struct {
	Identifier    string
	Name          string
	State         string
	Image         string
	Version       string
	CombinedImage string
	Create        time.Time

	Raw *types.ListContainer
}
