package image

import (
	"fmt"
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	monitorConstant "github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"time"
)

func New(v *types.ImageSummary) *Image {
	var first string

	if len(v.Names) == 0 {
		first = v.History[0]
	} else {
		first = v.Names[0]
	}

	name, version := key_value.Colon(first)
	t := time.Unix(v.Created, 0)

	return &Image{
		MonitorIdentifier: fmt.Sprintf(
			"%s-%s",
			monitorConstant.PodManPrefix,
			v.ID[:12],
		),
		Identifier:   v.ID,
		Name:         name,
		Version:      version,
		CombinedName: first,
		Dangling:     v.Dangling,
		Create:       &t,
		Raw:          v,
	}
}
