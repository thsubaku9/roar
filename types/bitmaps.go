package roar

//bitmaps stores 2^16 values in 32bit words -> 1024 entries
type Bitmaps struct {
	Values []uint32
	CType  containerType
}

func CreateBitmap() Bitmaps {
	return Bitmaps{
		Values: make([]uint32, 1024),
		CType:  bmps,
	}
}

func (bmp *Bitmaps) Add(elem uint16) error {
	index := elem / 32
	offset := elem % 32
	bmp.Values[index] |= (1 << offset)
	return nil
}

//func (bmp *Bitmaps) Union()
