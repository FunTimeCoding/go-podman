package image

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestImage(t *testing.T) {
	assert.True(
		t,
		New(&types.ImageSummary{Names: []string{"test:latest"}}) != nil,
	)
}
