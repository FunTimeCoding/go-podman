package container

import (
	"github.com/containers/podman/v5/libpod/define"
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"time"
)

type Container struct {
	MonitorIdentifier string
	Identifier        string
	Name              string
	State             string
	Image             string
	Version           string
	CombinedImage     string
	Restart           string
	Create            *time.Time

	Concern []string

	RawList   *types.ListContainer
	RawDetail *define.InspectContainerData
}
