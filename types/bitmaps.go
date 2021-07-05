package roar

import (
	"fmt"
	"math/bits"
	"roar/util"
)

//bitmaps stores 2^16 values in 32bit words -> 2048 entries
type Bitmaps struct {
	Values []uint32
	CType  util.ContainerType
}

func CreateBitmap() Bitmaps {
	return Bitmaps{
		Values: make([]uint32, util.BmpsLen),
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
	var shiftPos int = util.BmpRange - 1
	for i := util.BmpsLen - 1; i >= 0; i-- {
		if bmp.Values[i] != 0 {
			for ; bmp.Values[i]&(0x01<<shiftPos) == 0x00 && shiftPos >= 0; shiftPos-- {
			}
			offset := 32 * i
			return uint16(offset + shiftPos), nil
		}
	}
	return 0, fmt.Errorf("EmptyBitmapError")
}

func (bmp *Bitmaps) Min() (uint16, error) {
	var shiftPos uint32 = 0
	for i := 0; i < util.BmpsLen; i++ {
		if bmp.Values[i] != 0 {
			for ; bmp.Values[i]&(0x01<<shiftPos) == 0x00 && shiftPos < uint32(util.BmpsLen); shiftPos++ {
			}
			offset := uint32(32 * i)
			return uint16(offset + shiftPos), nil
		}
	}
	return 0, fmt.Errorf("EmptyBitmapError")
}

//Pop removes the element with highest value
func (bmp *Bitmaps) Pop() (uint16, error) {
	_max, err := bmp.Max()
	bmp.Remove(_max)
	return _max, err
}

//Select returns the element at the i-th index
func (bmp *Bitmaps) Select(index uint16) (uint16, error) {
	totalElems := uint16(0)

	for i, v := range bmp.Values {
		if v != 0 {
			for j := 0; j < util.BmpRange; j++ {
				if v&(0x01<<j) != 0x00 {
					if totalElems == index {
						offset := i * 32
						return uint16(offset + j), nil
					}
					totalElems++
				}
			}
		}
	}
	return 0, fmt.Errorf("IndexOutOfBounds")
}

//Index returns the index location of provided element
func (bmp *Bitmaps) Index(elem uint16) (uint16, error) {
	totalElems := uint16(0)

	for i, v := range bmp.Values {
		if v != 0 {
			for j := 0; j < util.BmpRange; j++ {
				if v&(0x01<<j) != 0x00 {
					offset := i * 32
					if uint16(offset+j) == elem {
						return totalElems, nil
					}
					totalElems++
				}
			}
		}
	}
	return 0, fmt.Errorf("ElementNotFound")
}

//Rank returns number of elements -le the given number
func (bmp *Bitmaps) Rank(elem uint16) uint16 {
	var currentCount uint16
	for i, v := range bmp.Values {
		if v != 0 {
			for j := 0; j < util.BmpRange; j++ {
				if v&(0x01<<j) != 0x00 {
					if uint16(i*32+j) > elem {
						return currentCount
					}
					currentCount++
				}
			}
		}
	}
	return currentCount
}

func (bmp *Bitmaps) NumElem() uint16 {
	totalElems := 0
	for _, v := range bmp.Values {
		totalElems += bits.OnesCount32(v)
	}
	return uint16(totalElems)
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
	_bmp := CreateBitmap()

	for i := 0; i < util.BmpsLen; i++ {
		_bmp.Values[i] = bmp.Values[i] ^ bmp2.Values[i]
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
			for ; v&(1<<iter) == 0; iter++ {
				if iter >= util.BmpRange {
					break innerL
				}
			}

			_start = iter
			for ; v&(1<<iter) > 0; iter++ {
			}
			_end = iter - 1

			_rles.RlePairs = append(_rles.RlePairs, RlePair{uint16(offset + _start), uint16(_end - _start)})
		}
	}
	//compact the _rles array
	_rles.compact()
	return _rles
}
