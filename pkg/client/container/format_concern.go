package container

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (c *Container) formatConcern(f *option.Format) string {
	if !c.HasConcerns() {
		return ""
	}

	result := join.Comma(c.Concerns())

	if f.UseColor {
		result = console.Yellow("%s", result)
	}

	return result
}
