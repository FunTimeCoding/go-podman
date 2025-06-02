package log

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"strings"
)

func FromStrings(
	standardOutput string,
	standardError string,
) *Log {
	l := New()

	for _, i := range split.NewLine(standardOutput) {
		if i == separator.Unix {
			continue
		}

		l.Output.Add(strings.TrimSpace(i))
	}

	for _, i := range split.NewLine(standardError) {
		if i == separator.Unix {
			continue
		}

		l.Error.Add(strings.TrimSpace(i))
	}

	return l
}
