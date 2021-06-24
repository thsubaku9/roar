package roar

import "roar/util"

type rle_pair struct {
	Start  uint16
	RunLen uint16
}
type Rles struct {
	RlePair []rle_pair
	CType   util.ContainerType
}

func CreateRles() Rles {
	return Rles{make([]rle_pair, 0), util.Rles}
}
