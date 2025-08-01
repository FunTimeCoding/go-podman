package container

import (
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/report"
	"github.com/funtimecoding/go-podman/pkg/check/container/option"
	"github.com/funtimecoding/go-podman/pkg/client/container"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func printNotation(
	v []*container.Container,
	o *option.Container,
) {
	r := report.New()
	f := constant.Format

	for _, e := range report.Trim(
		v,
		r,
		o.All,
		item.GoContainer,
	) {
		var s monitor.Severity

		if e.HasConcerns() {
			s = monitor.Warning
		} else {
			s = monitor.Information
		}

		r.AddItem(
			item.GoContainer,
			e.MonitorIdentifier,
			s,
			e.Format(f),
			"",
			e.Create,
		)
	}

	r.Print()
}
