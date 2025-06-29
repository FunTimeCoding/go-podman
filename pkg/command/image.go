package command

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/funtimecoding/go-podman/pkg/command/image"
	"github.com/funtimecoding/go-podman/pkg/podman"
)

func Images() []*image.Image {
	r := run.New()
	r.Start(
		podman.Command,
		podman.Image,
		podman.List,
		podman.Format,
		podman.Notation,
	)
	var images []*image.Image
	notation.DecodeStrict(r.OutputString, &images, true)

	return images
}
