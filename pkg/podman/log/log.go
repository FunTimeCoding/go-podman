package log

import "github.com/funtimecoding/go-library/pkg/text/multi_line"

type Log struct {
	Output *multi_line.MultiLine
	Error  *multi_line.MultiLine
}
