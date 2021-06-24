package roar

import (
	"fmt"
	"roar/util"
)

type Sarr struct {
	Arr   []uint16
	CType util.ContainerType
}

//overhead on user to sort and provide these values
func CreateSarr(val ...uint16) Sarr {
	return Sarr{val, util.Sarr}
}

//findIndex finds the location after which the given elem should be inserted
func (ar *Sarr) findIndex(elem uint16, start, end int) (int, error) {
	if end < start {
		return -1, fmt.Errorf("array is empty")
	}

	mid := (start + end) / 2

	if start == end {
		if ar.Arr[mid] == elem {
			return mid, nil
		}
		if ar.Arr[mid] < elem {
			return mid, fmt.Errorf("value doesn't exist")
		} else {
			return mid - 1, fmt.Errorf("value doesn't exist")
		}
	}
	if ar.Arr[mid] < elem {
		return ar.findIndex(elem, mid+1, end)
	} else {
		return ar.findIndex(elem, start, mid)
	}
}

func (ar *Sarr) Add(elem uint16) {
	if len(ar.Arr) > 0 {
		if elem > ar.Arr[len(ar.Arr)-1] {
			ar.Arr = append(ar.Arr, elem)
		} else if elem < ar.Arr[0] {
			ar.Arr = append([]uint16{elem}, ar.Arr...)
		} else if index, err := ar.findIndex(elem, 0, len(ar.Arr)-1); err != nil {
			index += 1
			var newArr []uint16
			newArr = append(newArr, ar.Arr[0:index]...)
			newArr = append(newArr, elem)
			newArr = append(newArr, ar.Arr[index:len(ar.Arr)]...)
			ar.Arr = newArr
		}
	} else {
		ar.Arr = append(ar.Arr, elem)
	}
}

func (ar *Sarr) Remove(elem uint16) {
	if len(ar.Arr) > 0 {
		if elem == ar.Arr[len(ar.Arr)-1] {
			ar.Arr = ar.Arr[0 : len(ar.Arr)-1]
		} else if elem == ar.Arr[0] {
			ar.Arr = ar.Arr[1:len(ar.Arr)]
		} else if index, err := ar.findIndex(elem, 0, len(ar.Arr)-1); err == nil {
			var newArr []uint16
			newArr = append(newArr, ar.Arr[0:index]...)
			newArr = append(newArr, ar.Arr[index+1:len(ar.Arr)]...)
			ar.Arr = newArr
		}
	}
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

	for i, j := 0, 0; i < len(ar.Arr) && j < len(ar2.Arr); i, j = i+1, j+1 {
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
