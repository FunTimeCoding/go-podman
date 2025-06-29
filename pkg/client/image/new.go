package image

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"time"
)

func New(v *types.ImageSummary) *Image {
	name, version := key_value.Colon(v.Names[0])

	return &Image{
		Identifier:   v.ID,
		Name:         name,
		Version:      version,
		CombinedName: v.Names[0],
		Create:       time.Unix(v.Created, 0),
		Raw:          v,
	}
}
