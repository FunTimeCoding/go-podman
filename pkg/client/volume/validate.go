package volume

import "slices"

func (v *Volume) Validate() {
	if v.RawDetail != nil &&
		v.Bytes == 0 &&
		!slices.Contains(v.concern, ZeroBytes) {
		v.concern = append(v.concern, ZeroBytes)
	}
}
