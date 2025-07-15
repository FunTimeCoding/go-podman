package container

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (c *Container) Format(f *option.Format) string {
	return status.New(f).String(
		c.formatName(f),
		c.State,
		c.Image,
		c.Version,
		c.Create.Format(time.DateMinute),
		c.formatConcern(f),
	).RawList(c.RawList).RawDetail(c.RawDetail).Format()
}
