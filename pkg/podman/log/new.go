package log

import "github.com/funtimecoding/go-library/pkg/text/multi_line"

func New() *Log {
	return &Log{Output: multi_line.New(), Error: multi_line.New()}
}
