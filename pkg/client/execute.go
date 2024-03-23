package client

import (
	"bytes"
	"github.com/containers/podman/v5/pkg/bindings/containers"
	"github.com/containers/podman/v5/pkg/domain/entities"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/log"
)

func (c *Client) Execute(
	v entities.ListContainer,
	command ...string,
) *log.Log {
	outputBuffer := &bytes.Buffer{}
	errorBuffer := &bytes.Buffer{}
	o := new(containers.ExecStartAndAttachOptions)
	o.WithOutputStream(outputBuffer)
	o.WithErrorStream(errorBuffer)
	o.WithAttachOutput(true)
	o.WithAttachError(true)
	session := c.session(v.ID, command...)
	errors.PanicOnError(containers.ExecStartAndAttach(c.context, session, o))
	c.wait(session)

	return log.FromStrings(outputBuffer.String(), errorBuffer.String())
}
