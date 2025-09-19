package podman

import "github.com/funtimecoding/go-library/pkg/system/environment"

func ContainerAutoStart() bool {
	return environment.Default(
		ContainerAutoStartEnvironment,
		"",
	) != ""
}
