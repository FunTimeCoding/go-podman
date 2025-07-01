package image

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (i *Image) Format(f *option.Format) string {
	return status.New(f).String(
		i.Name,
		i.Version,
		i.Create.Format(time.DateMinute),
		i.formatConcern(f),
	).RawList(i.Raw).Format()
}
