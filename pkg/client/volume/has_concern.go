package volume

import "slices"

func (v *Volume) HasConcern(s string) bool {
	return slices.Contains(v.concern, s)
}
