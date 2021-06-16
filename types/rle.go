package roar

import "roar/util"

type Rle struct {
	start  uint16
	runLen uint16
	CType  util.ContainerType
}
