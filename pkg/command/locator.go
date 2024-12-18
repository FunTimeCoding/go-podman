package command

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-podman/pkg/podman"
	"os"
)

func Locator() string {
	if locator := os.Getenv(podman.LocatorEnvironment); locator != "" {
		return locator
	}

	var result []Connection
	notation.DecodeStrict(
		system.Run(
			podman.Command,
			podman.System,
			podman.Connection,
			podman.List,
			podman.Format,
			podman.Notation,
		),
		&result,
		false,
	)

	return fmt.Sprintf("%s?secure=true", result[0].URI)
}
