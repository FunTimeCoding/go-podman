package client

import (
	"context"
	"github.com/containers/podman/v5/pkg/bindings"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-podman/pkg/command"
)

func New() *Client {
	result, e := bindings.NewConnectionWithIdentity(
		context.Background(),
		command.Locator(),
		command.Identity(),
		true,
	)
	errors.PanicOnError(e)

	return &Client{context: result}
}
