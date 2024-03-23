package container

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestContainer(t *testing.T) {
	actual := New(
		"testName",
		"testAlias",
		1,
		map[string]string{"testKey": "testValue"},
	)
	assert.String(t, "testName", actual.Name)
	assert.String(t, "testAlias", actual.Alias)
	assert.Integer(t, 1, actual.Port)
	assert.String(t, "testValue", actual.Environment["testKey"])
}
