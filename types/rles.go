package roar

import (
	"fmt"
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
func (rle *Rles) findIndex(p rlePair, start, end int) (int, error) {
	if end < start {
		return -1, fmt.Errorf("array is empty")
	}

	mid := (start + end) / 2

	if start == end {
		//check overlap region
		if rle.RlePairs[mid].isSubSegment(p) || rle.RlePairs[mid].lSideOverlap(p) || rle.RlePairs[mid].rSideOverlap(p) || p.isSubSegment(rle.RlePairs[mid]) {
			return mid, nil
		}
		if p.Start+p.RunLen < rle.RlePairs[mid].Start {
			return mid - 1, fmt.Errorf("no overlap")
		}
		if rle.RlePairs[mid].Start+rle.RlePairs[mid].RunLen < p.Start {
			return mid, fmt.Errorf("no overlap")
		}
	}

	if rle.RlePairs[mid].Start+rle.RlePairs[mid].RunLen < p.Start {
		return rle.findIndex(p, mid+1, end)
	}

	//p.Start+p.RunLen <= rle.RlePairs[mid].Start
	return rle.findIndex(p, start, mid)

}

func (p1 rlePair) lSideOverlap(p2 rlePair) bool {
	return (p2.Start+p2.RunLen >= p1.Start) && p2.Start+p2.RunLen <= p1.Start+p1.RunLen
}

func (p1 rlePair) rSideOverlap(p2 rlePair) bool {
	return (p2.Start <= p1.Start+p1.RunLen) && (p1.Start <= p2.Start)
}

func (p1 rlePair) isSubSegment(p2 rlePair) bool {
	return p1.Start <= p2.Start && p1.Start+p1.RunLen >= p2.Start+p2.RunLen
}

func (p1 rlePair) mergeReturn(p2 rlePair) rlePair {
	return rlePair{p1.Start, p1.RunLen + p2.RunLen - (p1.Start + p1.RunLen - p2.Start)}
}

//canMerge checks not overlap, but successive sequences for given pairs
func (p1 rlePair) canMerge(p2 rlePair) bool {
	return p2.Start-p1.Start+p1.RunLen == 1 || p1.Start-p2.Start+p2.RunLen == 1
}

func (p1 rlePair) splitReturn(p2 rlePair) []rlePair {
	//TODO -> implement splitReturn for Remove func
	return []rlePair{p1}
}

func (rle *Rles) Add(p rlePair) {
	if len(rle.RlePairs) == 0 {
		rle.RlePairs = append(rle.RlePairs, p)
		return
	}
	pos, _ := rle.findIndex(p, 0, len(rle.RlePairs)-1)
	pos = pos + 1

	_rles := rle.RlePairs[0:pos]
	_rles_l_index := len(_rles) - 1

	//check successive overlaps
	if _rles_l_index >= 0 {
		if _rles[_rles_l_index].lSideOverlap(p) {

		} else if _rles[_rles_l_index].rSideOverlap(p) {

		} else if _rles[_rles_l_index].isSubSegment(p) {

		} else {
			_rles = append(_rles, p)
		}
	} else {
		_rles = append(_rles, p)
	}

	for _, v := range rle.RlePairs[pos:len(rle.RlePairs)] {

	}

}

func (rle *Rles) Remove(p rlePair) {
	if len(rle.RlePairs) == 0 {
		return
	}
	pos, _ := rle.findIndex(p, 0, len(rle.RlePairs)-1)
}

func (rle *Rles) Union(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}

func (rle *Rles) Intersection(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}
