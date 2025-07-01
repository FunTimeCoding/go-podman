package image

import (
	monitorConstant "github.com/funtimecoding/go-library/pkg/monitor/constant"
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

	for _, n := range report.Trim(
		v,
		r,
		o.All,
		"images",
		monitorConstant.PodManPrefix,
	) {
		r.AddItem(
			n.MonitorIdentifier,
			monitorConstant.WarningLevel,
			n.Format(f),
			"",
			n.Create,
		)
	}

	r.Print()
}
