package podman

import (
	"context"
	"github.com/containers/podman/v4/pkg/bindings/containers"
	"github.com/containers/podman/v4/pkg/domain/entities"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/podman/log"
	"strings"
)

func (c *Client) Logs(o entities.ListContainer) *log.Log {
	var e error
	outputChannel := make(chan string)
	errorChannel := make(chan string)
	cancelContext, cancel := context.WithCancel(context.Background())
	go func() {
		e = containers.Logs(
			c.context,
			o.ID,
			new(containers.LogOptions).WithStdout(true).WithStderr(true),
			outputChannel,
			errorChannel,
		)
		cancel()
	}()
	l := log.New()

	for {
		select {
		case <-cancelContext.Done():
			errors.PanicOnError(e)

			return l
		case s := <-outputChannel:
			l.Output.Add(strings.TrimSpace(s))
		case s := <-errorChannel:
			l.Error.Add(strings.TrimSpace(s))
		}
	}
}
