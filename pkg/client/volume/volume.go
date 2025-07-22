package volume

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"time"
)

type Volume struct {
	MonitorIdentifier string
	Name              string
	Create            *time.Time

	Bytes   int64
	concern []string

	Raw       *types.VolumeListReport
	RawDetail *types.VolumeConfigResponse
}
