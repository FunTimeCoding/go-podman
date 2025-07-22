package image

func (i *Image) HasConcerns() bool {
	return len(i.concern) > 0
}
