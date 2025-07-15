package image

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Image) formatName(f *option.Format) string {
	result := i.Name

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
