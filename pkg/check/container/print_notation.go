package container

import (
	monitor "github.com/funtimecoding/go-library/pkg/monitor/constant"
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
		Plural,
		monitor.PodManPrefix,
	) {
		r.AddItem(
			e.MonitorIdentifier,
			monitor.WarningLevel,
			e.Format(f),
			"",
			e.Create,
		)
	}

	r.Print()
}
