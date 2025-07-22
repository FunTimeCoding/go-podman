package volume

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"time"
)

type Volume struct {
	Identifier string
	Name       string
	Create     *time.Time
	Raw        *types.VolumeListReport
}
