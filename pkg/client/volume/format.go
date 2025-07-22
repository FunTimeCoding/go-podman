package volume

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (v *Volume) Format(f *option.Format) string {
	return status.New(f).String(
		v.formatName(f),
		v.Create.Format(time.DateMinute),
		v.formatSize(f),
		v.formatConcern(f),
	).RawList(v.Raw).RawDetail(v.RawDetail).Format()
}
