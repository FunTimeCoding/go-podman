package volume

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor"
	"github.com/funtimecoding/go-podman/pkg/check/volume/option"
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/constant"
)

func Check(o *option.Volume) {
	c := client.New()
	elements := monitor.OnlyConcerns(c.Volume(true), o.All)

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(Plural)
	}
}
