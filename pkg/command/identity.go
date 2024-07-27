package command

import (
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/system"
	"strings"
)

func Identity() string {
	var machines []Machine
	notation.DecodeStrict(
		system.Run("podman", "machine", "inspect"),
		&machines,
		false,
	)

	result := machines[0].SSHConfig.IdentityPath

	if strings.HasPrefix(result, "C:") {
		result = system.WindowsToLinuxPath(result)
		result = strings.TrimPrefix(result, "/c")
	}

	return result
}
