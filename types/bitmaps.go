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

func (bmp *Bitmaps) Remove(elem uint16) error {
	index := elem / 32
	offset := elem % 32
	bmp.Values[index] &= 0xFFFF ^ (1 << offset)
	return nil
}

/*
Clamp(start, stop uint32)
Clear()
Copy() Container
Debug() string
FlipRange(start, stop uint32)
Index(element uint32) (uint32, error) //returns the index location of provided element
IsSuperset(con Container) bool
Jaccard(con Container) float32
Max() (uint32, error)
Min() (uint32, error)
Pop() (uint32, error)         //removes the element with highest value
Rank(element uint32) []uint32 //number of elements -le the given number
Select(index uint32) (uint32, error) //return the element at the i-th index
SymmetricDifference(con Container) (Container, error)
*/

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

func (bmp *Bitmaps) Intersection(bmp2 *Bitmaps) Bitmaps {
	_bmp := Bitmaps{
		Values: make([]uint32, 1024),
		CType:  util.Bmps,
	}

	for i := range bmp.Values {
		_bmp.Values[i] = (*bmp).Values[i] & (*bmp2).Values[i]
	}
	return _bmp
}

func (bmp *Bitmaps) Difference(bmp2 *Bitmaps) Bitmaps {
	_bmp := Bitmaps{
		Values: make([]uint32, 1024),
		CType:  util.Bmps,
	}
	for i := range bmp.Values {
		_bmp.Values[i] = (*bmp).Values[i] ^ ((*bmp).Values[i] & (*bmp2).Values[i])
	}
	return _bmp
}

func (bmp *Bitmaps) IsDisjoint(bmp2 *Bitmaps) bool {
	for i := range bmp.Values {
		if (*bmp).Values[i]&(*bmp2).Values[i] != 0 {
			return false
		}
	}
	return true
}

func (bmp *Bitmaps) IsSubset(bmp2 *Bitmaps) bool {
	return bmp2.IsSuperset(bmp)
}

func (bmp *Bitmaps) IsSuperset(bmp2 *Bitmaps) bool {
	for i := range bmp.Values {
		if (*bmp).Values[i]|(*bmp2).Values[i] != (*bmp).Values[i] {
			return false
		}
	}
	return true
}
