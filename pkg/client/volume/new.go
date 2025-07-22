package volume

import "github.com/containers/podman/v5/pkg/domain/entities/types"

func New(v *types.VolumeListReport) *Volume {
	return &Volume{
		//MonitorIdentifier: fmt.Sprintf(
		//	"%s-%s",
		//	monitorConstant.PodManPrefix,
		//	v.ID[:12],
		//),
		Identifier: v.StorageID,
		Name:       v.Name,
		//Version:      version,
		//CombinedName: first,
		//Dangling:     v.Dangling,
		Create: &v.CreatedAt,
		Raw:    v,
	}
}
