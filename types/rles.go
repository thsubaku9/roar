package roar

import (
	"roar/util"
)

type rlePair struct {
	Start  uint16
	RunLen uint16
}
type Rles struct {
	RlePairs []rlePair
	CType    util.ContainerType
}

func CreateRles() Rles {
	return Rles{make([]rlePair, 0), util.Rles}
}

//findIndex finds the location after which the given elem should be inserted
/*
func (rle *Rles) findIndex(p rlePair, start, end int) (int, error) {
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
*/

//only does increasing order check
func (p1 rlePair) checkOverlap(p2 rlePair) bool {
	return p1.Start+p1.RunLen >= p2.Start
}

func (p1 rlePair) mergeReturn(p2 rlePair) rlePair {
	return rlePair{p1.Start, p1.RunLen + p2.RunLen - (p1.Start + p1.RunLen - p2.Start)}
}

func (rle *Rles) Add(p rlePair) {
	if len(rle.RlePairs) == 0 {
		rle.RlePairs = append(rle.RlePairs, p)
		return
	}
}

func (rle *Rles) Remove(p rlePair) {
	if len(rle.RlePairs) == 0 {
		return
	}
}

func (rle *Rles) Union(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}

func (rle *Rles) Intersection(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}
