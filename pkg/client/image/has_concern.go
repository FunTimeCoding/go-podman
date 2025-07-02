package image

import "slices"

func (i *Image) HasConcern(s string) bool {
	return slices.Contains(i.Concern, s)
}
