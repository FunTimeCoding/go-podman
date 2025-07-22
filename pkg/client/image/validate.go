package image

import "slices"

func (i *Image) Validate() {
	if i.ShouldDelete() && !slices.Contains(i.concern, ShouldDelete) {
		i.concern = append(i.concern, ShouldDelete)
	}
}
