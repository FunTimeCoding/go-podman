package volume

import "github.com/containers/podman/v5/pkg/domain/entities/types"

func (v *Volume) Enrich(o *types.VolumeConfigResponse) {
	v.RawDetail = o
	v.Validate()
}
