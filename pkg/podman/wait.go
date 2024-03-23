package podman

import (
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/funtimecoding/go-library/pkg/errors"
	"time"
)

func (c *Client) wait(session string) {
	for {
		inspect, e := containers.ExecInspect(
			c.context,
			session,
			&containers.ExecInspectOptions{},
		)
		errors.PanicOnError(e)

		if !inspect.Running {
			break
		}

		time.Sleep(time.Second)
	}
}
