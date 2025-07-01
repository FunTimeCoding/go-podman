package image

import (
	"github.com/funtimecoding/go-podman/pkg/client"
	"github.com/funtimecoding/go-podman/pkg/podman"
)

func autoUpdate(c *client.Client) {
	if !podman.ImageAutoUpdate() {
		return
	}

	// TODO: Checks all images for updates
	//  Easy if version is latest or main
	//  SemVer tags also, but what if not wanting to take next major or minor?
	//  And then there are git hash tags and various other conventions
}
