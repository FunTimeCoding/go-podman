package volume

import "github.com/containers/podman/v5/pkg/domain/entities/types"

func NewSlice(v []*types.VolumeListReport) []*Volume {
	var result []*Volume

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
