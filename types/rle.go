package roar

import "roar/util"

type Rle struct {
	Start  uint16
	RunLen uint16
	CType  util.ContainerType
}

func CreateRle(start, runlen uint16) Rle {
	return Rle{start, runlen, util.Rles}
}
