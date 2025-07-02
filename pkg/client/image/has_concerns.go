package image

func (i *Image) HasConcerns() bool {
	return len(i.Concern) > 0
}
