package roar

import "roar/util"

//bitmaps stores 2^16 values in 32bit words -> 1024 entries
type Bitmaps struct {
	Values []uint32
	CType  util.ContainerType
}

func CreateBitmap() Bitmaps {
	return Bitmaps{
		Values: make([]uint32, 1024),
		CType:  util.Bmps,
	}
}

func (bmp *Bitmaps) Add(elem uint16) error {
	index := elem / 32
	offset := elem % 32
	bmp.Values[index] |= (1 << offset)
	return nil
}

func (bmp *Bitmaps) Union(bmp2 *Bitmaps) Bitmaps {
	_bmp := Bitmaps{
		Values: make([]uint32, 1024),
		CType:  util.Bmps,
	}

	for i := range bmp.Values {
		_bmp.Values[i] = (*bmp).Values[i] | (*bmp2).Values[i]
	}
	return _bmp
}
