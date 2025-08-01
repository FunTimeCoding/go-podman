package volume

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"github.com/funtimecoding/go-library/pkg/monitor/item/constant"
)

func New(v *types.VolumeListReport) *Volume {
	return &Volume{
		MonitorIdentifier: constant.GoVolume.StringIdentifier(v.Name),

		Name:   v.Name,
		Create: &v.CreatedAt,
		Raw:    v,
	}
}
