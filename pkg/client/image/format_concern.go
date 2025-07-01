package image

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (i *Image) formatConcern(f *option.Format) string {
	if len(i.Concern) == 0 {
		return ""
	}

	result := join.Comma(i.Concern)

	if f.UseColor {
		result = console.Yellow("%s", result)
	}

	return result
}
