package volume

import (
	"fmt"
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	monitorConstant "github.com/funtimecoding/go-library/pkg/monitor/constant"
)

func New(v *types.VolumeListReport) *Volume {
	return &Volume{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%s",
			monitorConstant.PodManPrefix,
			v.Name,
		),
		Name:   v.Name,
		Create: &v.CreatedAt,
		Raw:    v,
	}
}
