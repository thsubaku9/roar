package roar

import (
	"fmt"
	"roar/util"
	"sort"
)

type Sarr struct {
	Arr   []uint16
	CType util.ContainerType
}

//CreateSarr takes an array of uint16 irrsepective of being sorted
func CreateSarr(val ...uint16) Sarr {
	if sort.SliceIsSorted(val, func(i, j int) bool {
		return val[i] < val[j]
	}) {
		return Sarr{val, util.Sarr}
	}

	sort.SliceStable(val, func(i, j int) bool {
		return val[i] < val[j]
	})
	return Sarr{val, util.Sarr}
}

//findIndex finds the location after which the given elem should be inserted
func (ar *Sarr) findIndex(elem uint16, start, end int) (int, error) {
	if end < start {
		return -1, fmt.Errorf("EmptyArray")
	}

	mid := (start + end) / 2

	if start == end {
		if ar.Arr[mid] == elem {
			return mid, nil
		}
		if ar.Arr[mid] < elem {
			return mid, fmt.Errorf("ValueNotFound")
		} else {
			return mid - 1, fmt.Errorf("ValueNotFound")
		}
	}
	if ar.Arr[mid] < elem {
		return ar.findIndex(elem, mid+1, end)
	} else {
		return ar.findIndex(elem, start, mid)
	}
}

func (ar *Sarr) Add(elem uint16) {
	if len(ar.Arr) == 0 {
		ar.Arr = append(ar.Arr, elem)
		return
	}

	if elem > ar.Arr[len(ar.Arr)-1] {
		ar.Arr = append(ar.Arr, elem)
	} else if elem < ar.Arr[0] {
		ar.Arr = append([]uint16{elem}, ar.Arr...)
	} else if index, err := ar.findIndex(elem, 0, len(ar.Arr)-1); err != nil {
		index += 1
		var newArr []uint16
		newArr = append(newArr, ar.Arr[:index]...)
		newArr = append(newArr, elem)
		newArr = append(newArr, ar.Arr[index:]...)
		ar.Arr = newArr
	}
}

func (ar *Sarr) Remove(elem uint16) {
	if len(ar.Arr) == 0 {
		return
	}

	if elem == ar.Arr[len(ar.Arr)-1] {
		ar.Arr = ar.Arr[:len(ar.Arr)-1]
	} else if elem == ar.Arr[0] {
		ar.Arr = ar.Arr[1:]
	} else if index, err := ar.findIndex(elem, 0, len(ar.Arr)-1); err == nil {
		var newArr []uint16
		newArr = append(newArr, ar.Arr[:index]...)
		newArr = append(newArr, ar.Arr[index+1:]...)
		ar.Arr = newArr
	}
}

func (ar *Sarr) Max() (uint16, error) {
	if len(ar.Arr) == 0 {
		return 0, fmt.Errorf("EmtpySarrError")
	}
	return ar.Arr[len(ar.Arr)-1], nil
}

func (ar *Sarr) Min() (uint16, error) {
	if len(ar.Arr) == 0 {
		return 0, fmt.Errorf("EmtpySarrError")
	}
	return ar.Arr[0], nil
}

func (ar *Sarr) NumElem() uint16 {
	return uint16(len(ar.Arr))
}

func (ar *Sarr) Pop() (uint16, error) {
	elem, err := ar.Max()
	if err != nil {
		ar.Remove(elem)
	}
	return elem, err
}

//Select returns the element at the i-th index
func (ar *Sarr) Select(index uint16) (uint16, error) {
	if index < uint16(len(ar.Arr)) {
		return ar.Arr[index], nil
	}
	return 0, fmt.Errorf("IndexOutOfBounds")
}

//Index returns the index location of provided element
func (ar *Sarr) Index(elem uint16) (int, error) {
	return ar.findIndex(elem, 0, len(ar.Arr)-1)
}

//Rank returns number of elements -le the given number
func (ar *Sarr) Rank(elem uint16) uint16 {
	index, _ := ar.findIndex(elem, 0, len(ar.Arr)-1)
	return uint16(index + 1)
}

func (ar *Sarr) Union(ar2 *Sarr) Sarr {
	_retSarr := CreateSarr()
	var i, j int
	for i < len(ar.Arr) && j < len(ar2.Arr) {
		switch {
		case ar.Arr[i] < ar2.Arr[j]:
			_retSarr.Arr = append(_retSarr.Arr, ar.Arr[i])
			i++
		case ar.Arr[i] > ar2.Arr[j]:
			_retSarr.Arr = append(_retSarr.Arr, ar2.Arr[j])
			j++
		default:
			_retSarr.Arr = append(_retSarr.Arr, ar.Arr[i])
			i++
			j++
		}
	}

	for i < len(ar.Arr) {
		_retSarr.Arr = append(_retSarr.Arr, ar.Arr[i])
		i++
	}
	for j < len(ar2.Arr) {
		_retSarr.Arr = append(_retSarr.Arr, ar2.Arr[j])
		j++
	}

	return _retSarr
}

func (ar *Sarr) Intersection(ar2 *Sarr) Sarr {
	_retSarr := Sarr{}

	for i, j := 0, 0; i < len(ar.Arr) && j < len(ar2.Arr); {
		switch {
		case ar.Arr[i] < ar2.Arr[j]:
			i++
		case ar.Arr[i] > ar2.Arr[j]:
			j++
		default:
			_retSarr.Arr = append(_retSarr.Arr, ar.Arr[i])
			i++
			j++
		}
	}

	return _retSarr
}

func (ar *Sarr) Sarr2Bmps() Bitmaps {
	_bmp := CreateBitmap()
	for _, v := range ar.Arr {
		_bmp.Add(v)
	}
	return _bmp
}

func (ar *Sarr) Sarr2Rles() Rles {
	_rle := CreateRles()

	if len(ar.Arr) == 0 {
		return _rle
	}

	var startPos, endPos int
	for endPos = 1; endPos <= len(ar.Arr)-1; endPos++ {
		if ar.Arr[endPos-1] != ar.Arr[endPos]-1 {
			_rle.RlePairs = append(_rle.RlePairs, RlePair{Start: ar.Arr[startPos], RunLen: uint16(endPos - startPos - 1)})
			startPos = endPos
		}
	}

	_rle.RlePairs = append(_rle.RlePairs, RlePair{Start: ar.Arr[startPos], RunLen: uint16(endPos - startPos - 1)})
	_rle.compact()
	return _rle
}
