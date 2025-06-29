package container

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
)

func New(v *types.ListContainer) *Container {
	image, version := key_value.Colon(v.Image)

	return &Container{
		Identifier:    v.ID,
		Name:          v.Names[0],
		State:         v.State,
		Image:         image,
		Version:       version,
		CombinedImage: v.Image,
		Create:        v.Created,
		Raw:           v,
	}
}
