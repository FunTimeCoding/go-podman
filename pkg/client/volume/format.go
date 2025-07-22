package volume

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (v *Volume) Format(f *option.Format) string {
	return status.New(f).String(
		v.Identifier,
		v.Name,
		v.Create.Format(time.DateMinute),
	).RawList(v.Raw).Format()
}
