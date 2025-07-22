package volume

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (v *Volume) formatConcern(f *option.Format) string {
	if !v.HasConcerns() {
		return ""
	}

	result := join.Comma(v.Concerns())

	if f.UseColor {
		result = console.Yellow("%s", result)
	}

	return result
}
