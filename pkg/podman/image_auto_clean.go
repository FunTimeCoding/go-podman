package podman

import "github.com/funtimecoding/go-library/pkg/system/environment"

func ImageAutoClean() bool {
	return environment.GetDefault(
		ImageAutoCleanEnvironment,
		"",
	) != ""
}
