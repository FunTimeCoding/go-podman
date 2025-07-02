package image

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (i *Image) formatConcern(f *option.Format) string {
	if !i.HasConcerns() {
		return ""
	}

	result := join.Comma(i.Concerns())

	if f.UseColor {
		result = console.Yellow("%s", result)
	}

	return result
}
