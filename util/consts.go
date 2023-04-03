package util

type SubContainerType int

const (
	Bmps SubContainerType = iota
	Rles
	Sarr
)

const (
	Limit          int = 1 << 16
	BmpRange       int = 32
	BmpsLen        int = Limit / (BmpRange)
	PivotPointUp   int = 4096
	PivotPointDown int = 2048
)
