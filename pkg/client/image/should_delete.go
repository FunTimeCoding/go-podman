package image

func (i *Image) ShouldDelete() bool {
	return i.Dangling
}
