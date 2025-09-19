package podman

import "github.com/funtimecoding/go-library/pkg/system/environment"

func ImageAutoUpdate() bool {
	return environment.Default(
		ImageAutoUpdateEnvironment,
		"",
	) != ""
}
