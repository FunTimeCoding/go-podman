package command

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/integers"
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"net"
	"net/url"
	"runtime"
	"strings"
	"testing"
)

func replacePort(
	locator string,
	port int,
) string {
	result, parseFail := url.Parse(locator)
	errors.PanicOnError(parseFail)
	host, _, splitFail := net.SplitHostPort(result.Host)
	errors.PanicOnError(splitFail)

	if host == "localhost" {
		host = "127.0.0.1"
	}

	result.Host = net.JoinHostPort(host, integers.ToString(port))

	return result.String()
}

func TestCommand(t *testing.T) {
	switch runtime.GOOS {
	case constant.Darwin:
		assert.String(
			t,
			"ssh://core@127.0.0.1:2000/run/user/501/podman/podman.sock?secure=true",
			replacePort(Locator(), 2000),
		)
	default:
		assert.String(
			t,
			"ssh://user@127.0.0.1:2000/run/user/1000/podman/podman.sock?secure=true",
			replacePort(Locator(), 2000),
		)
	}

	assert.True(
		t,
		strings.HasSuffix(
			Identity(),
			".local/share/containers/podman/machine/machine",
		),
	)
}
