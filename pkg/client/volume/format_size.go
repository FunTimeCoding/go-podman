package volume

import (
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (v *Volume) formatSize(f *option.Format) string {
	var result string

	if v.RawDetail != nil {
		result = units.HumanSize(float64(v.Bytes))

		if f.UseColor {
			if v.Bytes == 0 {
				result = console.Red("%s", result)
			}
		}
	}

	return result
}
