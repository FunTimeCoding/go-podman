package volume

import (
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
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
		Plural,
		monitor.PodManPrefix,
	) {
		var level string

		if e.HasConcerns() {
			level = monitor.WarningLevel
		} else {
			level = monitor.InformationLevel
		}

		r.AddItem(
			e.MonitorIdentifier,
			level,
			e.Format(f),
			"",
			e.Create,
		)
	}

	r.Print()
}
