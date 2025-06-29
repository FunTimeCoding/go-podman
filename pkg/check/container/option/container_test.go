package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestContainer(t *testing.T) {
	assert.True(t, New() != nil)
}
