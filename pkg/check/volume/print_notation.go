package volume

import (
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-podman/pkg/check/volume/option"
	"github.com/funtimecoding/go-podman/pkg/client/volume"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func printNotation(
	v []*volume.Volume,
	o *option.Volume,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoVolume,
	) {
		var s monitor.Severity

		if e.HasConcerns() {
			s = monitor.Warning
		} else {
			s = monitor.Information
		}

		r.AddItem(
			item.GoVolume,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			e.Create,
		)
	}

	r.Print()
}
