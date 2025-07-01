package client

import (
	"fmt"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/bindings/images"
	"github.com/containers/podman/v5/pkg/domain/entities/reports"
	"github.com/containers/podman/v5/pkg/domain/entities/types"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Remove(name string) []*reports.RmReport {
	result, e := containers.Remove(
		c.context,
		name,
		&containers.RemoveOptions{},
	)
	errors.PanicOnError(e)

	return result
}

func (c *Client) RemoveImage(v []string) *types.ImageRemoveReport {
	result, e := images.Remove(
		c.context,
		v,
		&images.RemoveOptions{},
	)

	if len(e) > 0 {
		for _, r := range e {
			fmt.Printf("Error: %s\n", r.Error())
		}

		errors.PanicOnError(e[0])
	}

	return result
}
