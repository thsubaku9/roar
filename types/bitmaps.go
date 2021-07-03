package roar

import (
	"roar/util"
)

//bitmaps stores 2^16 values in 32bit words -> 2048 entries
type Bitmaps struct {
	Values []uint32
	CType  util.ContainerType
}

func CreateBitmap() Bitmaps {
	return Bitmaps{
		Values: make([]uint32, 2048),
		CType:  util.Bmps,
	}
}

func (bmp *Bitmaps) Add(elem uint16) {
	index := elem / 32
	offset := elem % 32
	bmp.Values[index] |= (1 << offset)
}

func (bmp *Bitmaps) Remove(elem uint16) {
	index := elem / 32
	offset := elem % 32
	bmp.Values[index] &= 0xFFFF ^ (1 << offset)
}

func (bmp *Bitmaps) Max() (uint16, error) {
	return 0, nil
}

func (bmp *Bitmaps) Min() (uint16, error) {
	return 0, nil
}

//Pop removes the element with highest value
func (bmp *Bitmaps) Pop() (uint32, error) {
	return 0, nil
}

//Select returns the element at the i-th index
func (bmp *Bitmaps) Select(index uint16) (uint16, error) {

}

//Index returns the index location of provided element
func (bmp *Bitmaps) Index(elem uint16) (uint16, error) {

}

//Rank returns number of elements -le the given number
func (bmp *Bitmaps) Rank(elem uint16) (uint16, error) {

}

func (bmp *Bitmaps) Union(bmp2 *Bitmaps) Bitmaps {
	_bmp := CreateBitmap()

	for i := range bmp.Values {
		_bmp.Values[i] = (*bmp).Values[i] | (*bmp2).Values[i]
	}
	return _bmp
}

func (bmp *Bitmaps) Intersection(bmp2 *Bitmaps) Bitmaps {
	_bmp := CreateBitmap()

	for i := range bmp.Values {
		_bmp.Values[i] = (*bmp).Values[i] & (*bmp2).Values[i]
	}
	return _bmp
}

func (bmp *Bitmaps) Difference(bmp2 *Bitmaps) Bitmaps {
	_bmp := CreateBitmap()
	for i := range bmp.Values {
		_bmp.Values[i] = (*bmp).Values[i] ^ ((*bmp).Values[i] & (*bmp2).Values[i])
	}
	return _bmp
}

//{1, 2, 3} - {0, 1} = {0, 2, 3}
func (bmp *Bitmaps) SymmetricDifference(bmp2 *Bitmaps) Bitmaps {

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

func (bmp *Bitmaps) Bmps2Sarr() Sarr {
	_sarr := CreateSarr()

	for i, v := range bmp.Values {
		offset := 32 * i
		for _v, k := v, 0; _v > 0; _v, k = _v>>1, k+1 {
			if _v&0x01 == 0x01 {
				_sarr.Add(uint16(offset + k))
			}
		}
	}
	return _sarr
}

func (bmp *Bitmaps) Bmps2Rles() Rles {
	_rles := CreateRles()
	for i, v := range bmp.Values {

		if v == 0 {
			continue
		}
		offset := 32 * i
		var iter, _start, _end int

	innerL:
		for iter < util.BmpRange {
			for ; (1 << iter & v) == 0; iter++ {
				if iter >= util.BmpRange {
					break innerL
				}
			}

			_start = iter
			for ; (1 << iter & v) > 0; iter++ {
			}
			_end = iter - 1

			_rles.RlePairs = append(_rles.RlePairs, RlePair{uint16(offset + _start), uint16(_end - _start)})
		}
	}
	//compact the _rles array
	return _rles
}
