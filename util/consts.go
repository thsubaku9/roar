package util

type SubContainerType int

const (
	Bmps SubContainerType = iota
	Rles
	Sarr
)

const (
	SplitVal uint32 = 65536
	BmpRange int    = 32
	BmpsLen  int    = 2048
)
