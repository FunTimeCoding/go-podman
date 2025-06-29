package container

import (
	monitorConstant "github.com/funtimecoding/go-library/pkg/monitor/constant"
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

	for _, n := range report.Trim(
		v,
		r,
		o.All,
		"containers",
		monitorConstant.PodManPrefix,
	) {
		r.AddItem(
			monitorConstant.PodManPrefix,
			monitorConstant.WarningLevel,
			n.Format(f),
			"",
			n.Create,
		)
	}

	r.Print()
}
