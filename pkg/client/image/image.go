package image

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"time"
)

type Image struct {
	Identifier   string
	Name         string
	Version      string
	CombinedName string
	Create       time.Time

	Raw *types.ImageSummary
}
