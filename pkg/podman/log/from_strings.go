package log

import (
	"github.com/funtimecoding/go-library/pkg/separator"
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"strings"
)

func FromStrings(
	standardOutput string,
	standardError string,
) *Log {
	l := New()

	for _, element := range split.NewLine(standardOutput) {
		if element == separator.Unix {
			continue
		}

		l.Output.Add(strings.TrimSpace(element))
	}

	for _, element := range split.NewLine(standardError) {
		if element == separator.Unix {
			continue
		}

		l.Error.Add(strings.TrimSpace(element))
	}

	return l
}
