package image

import (
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-podman/pkg/check/image/option"
	"github.com/funtimecoding/go-podman/pkg/client/image"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func printNotation(
	v []*image.Image,
	o *option.Image,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoImage,
	) {
		var s monitor.Severity

		if e.HasConcerns() {
			s = monitor.Warning
		} else {
			s = monitor.Information
		}

		r.AddItem(
			item.GoImage,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			e.Create,
		)
	}

	r.Print()
}
