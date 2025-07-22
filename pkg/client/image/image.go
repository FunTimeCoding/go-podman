package image

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"time"
)

type Image struct {
	MonitorIdentifier string
	Identifier        string
	Name              string
	Version           string
	CombinedName      string
	Dangling          bool
	Create            *time.Time

	concern []string

	Raw *types.ImageSummary
}
