package container

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (c *Container) formatName(f *option.Format) string {
	result := c.Name

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
