package image

import "github.com/funtimecoding/go-library/pkg/strings/split/key_value"

func (i *Image) NameVersion() (string, string) {
	return key_value.Colon(i.Name())
}
