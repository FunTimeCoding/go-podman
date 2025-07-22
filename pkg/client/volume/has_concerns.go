package volume

func (v *Volume) HasConcerns() bool {
	return len(v.concern) > 0
}
