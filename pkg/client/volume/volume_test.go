package volume

import (
	"github.com/containers/podman/v5/libpod/define"
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestVolume(t *testing.T) {
	assert.True(
		t,
		New(
			&types.VolumeListReport{
				VolumeConfigResponse: types.VolumeConfigResponse{
					InspectVolumeData: define.InspectVolumeData{
						Name: strings.Alfa,
					},
				},
			},
		) != nil,
	)
}
