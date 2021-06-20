package roar

import "roar/util"

type Sarr struct {
	Arr   []uint16
	CType util.ContainerType
}

//overhead on user to sort and provide these values
func CreateSarr(val ...uint16) Sarr {
	return Sarr{val, util.Sarr}
}
