package container

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (i *Container) Format(f *option.Format) string {
	return status.New(f).String(
		i.Name,
		i.State,
		i.Image,
		i.Version,
		i.Create.Format(time.DateMinute),
	).RawList(i.Raw).Format()
}
