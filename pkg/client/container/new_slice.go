package container

import "github.com/containers/podman/v5/pkg/domain/entities/types"

func NewSlice(v []types.ListContainer) []*Container {
	var result []*Container

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
