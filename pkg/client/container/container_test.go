package container

import (
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-podman/pkg/podman"
	"testing"
)

func TestImage(t *testing.T) {
	assert.True(
		t,
		New(
			&types.ListContainer{
				ID:    podman.FixtureIdentifier,
				Names: []string{podman.FixtureName},
			},
		) != nil,
	)
}
