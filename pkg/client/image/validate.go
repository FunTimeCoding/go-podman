package image

func (i *Image) Validate() {
	if i.ShouldDelete() {
		i.Concern = append(i.Concern, ShouldDelete)
	}
}
