package util

type ContainerType int

const (
	Bmps ContainerType = iota
	Rles
	Sarr
)

const (
	SplitVal uint32 = 65536
	BmpRange int    = 32
)
