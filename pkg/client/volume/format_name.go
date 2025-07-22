package volume

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (v *Volume) formatName(f *option.Format) string {
	result := v.Name

	if len(result) > 12 {
		result = result[:12]
	}

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
