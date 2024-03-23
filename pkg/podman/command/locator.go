package command

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
	"os"
)

func Locator() string {
	if locator := os.Getenv("PODMAN_LOCATOR"); locator != "" {
		return locator
	}

	var result []Connection
	notation.DecodeStrict(
		system.Run(
			"podman",
			"system",
			"connection",
			"ls",
			"--format",
			"json",
		),
		&result,
	)

	return fmt.Sprintf("%s?secure=true", result[0].URI)
}
