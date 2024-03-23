package client

import (
	"github.com/containers/podman/v4/pkg/bindings/pods"
	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Pod() []*entities.ListPodsReport {
	result, e := pods.List(c.context, &pods.ListOptions{})
	errors.PanicOnError(e)

	return result
}
