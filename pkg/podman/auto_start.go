package podman

import "github.com/funtimecoding/go-library/pkg/system/environment"

func AutoStart() bool {
	return environment.GetDefault(
		AutoStartEnvironment,
		"",
	) != ""
}
