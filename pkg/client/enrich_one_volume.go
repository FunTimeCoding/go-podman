package client

import (
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/strings/split/key_value"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/funtimecoding/go-podman/pkg/client/volume"
)

func (c *Client) enrichOneVolume(o *volume.Volume) {
	r := run.New()
	r.Start(
		"podman",
		"machine",
		"ssh",
		"--",
		"du",
		"-bs",
		o.Raw.Mountpoint,
	)
	bytes, _ := key_value.Tab(r.OutputString)
	o.Bytes = strings.ToInteger64Strict(bytes)
	o.Enrich(c.InspectVolume(o.Name))
}
