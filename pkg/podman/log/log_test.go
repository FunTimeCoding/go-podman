package log

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestLog(t *testing.T) {
	actual := FromStrings("a\nb", "")
	assert.Any(t, "a\nb", actual.Output.Format())
	assert.Any(t, "", actual.Error.Format())
}
