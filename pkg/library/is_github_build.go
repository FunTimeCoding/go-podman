package library

import "github.com/funtimecoding/go-library/pkg/system/environment"

func IsGitHubBuild() bool {
	return environment.Default("GITHUB_RUN_ID", "") != ""
}
