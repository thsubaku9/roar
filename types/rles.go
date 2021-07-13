package roar

import (
	"fmt"
	"roar/util"
)

type RlePair struct {
	Start  uint16
	RunLen uint16
}
type Rles struct {
	RlePairs []RlePair
	CType    util.ContainerType
}

func CreateRles() Rles {
	return Rles{make([]RlePair, 0), util.Rles}
}

func (p1 RlePair) lSideOverlap(p2 RlePair) bool {
	return (p2.Start+p2.RunLen >= p1.Start) && p2.Start+p2.RunLen <= p1.Start+p1.RunLen && p2.Start < p1.Start
}

func (p1 RlePair) rSideOverlap(p2 RlePair) bool {
	return (p2.Start >= p1.Start) && (p2.Start <= p1.Start+p1.RunLen) && p2.Start+p2.RunLen > p1.Start+p1.RunLen
}

func (p1 RlePair) isSubSegment(p2 RlePair) bool {
	return p1.Start <= p2.Start && p1.Start+p1.RunLen >= p2.Start+p2.RunLen
}

//overlapReturn assumes the two pairs do overlap and combines them
func (p1 RlePair) overlapReturn(p2 RlePair) RlePair {
	startP := Min(p1.Start, p2.Start)
	endP := Max(p1.Start+p1.RunLen, p2.Start+p2.RunLen)

	return RlePair{Start: startP, RunLen: endP - startP}
}

//intersectReturn assumes the two pairs overlap and provides intersection
func (p1 RlePair) intersectReturn(p2 RlePair) RlePair {
	startP := Min(p1.Start, p2.Start)
	endP := Min(p1.Start+p1.RunLen, p2.Start+p2.RunLen)
	return RlePair{Start: startP, RunLen: endP - startP}
}

//canMerge checks not overlap, but successive sequence for given pairs
func (p1 RlePair) canMerge(p2 RlePair) bool {
	return p2.Start-(p1.Start+p1.RunLen) == 1
}

//mergeReturn merges two disjoint pairs assuming canMerge holds true
func (p1 RlePair) mergeReturn(p2 RlePair) RlePair {
	//RlePair{minP.Start, minP.RunLen + maxP.RunLen - (minP.Start + minP.RunLen - maxP.Start)} is a superset of RlePair{minP.Start, minP.RunLen + maxP.RunLen + 1}
	return p1.overlapReturn(p2)
}

//splitReturn returns two disjoint pairs ({p1}/{p2})
func (p1 RlePair) splitReturn(p2 RlePair) (*RlePair, *RlePair) {
	if p1.Start >= p2.Start {
		if p1.Start+p1.RunLen <= p2.Start+p2.RunLen {
			//p1 is subsegment
			return nil, nil
		}
		//lSideOverlap
		return nil, &RlePair{p2.Start + p2.RunLen + 1, p1.Start + p1.RunLen - (p2.Start + p2.RunLen + 1)}
	} else if p1.Start < p2.Start {
		if p1.Start+p1.RunLen > p2.Start+p2.RunLen {
			//p2 is subsegment
			return &RlePair{p1.Start, p2.Start - 1 - p1.Start}, &RlePair{p2.Start + p2.RunLen + 1, p1.Start + p1.RunLen - (p2.Start + p2.RunLen + 1)}
		}
		//rSideOverlap
		return &RlePair{p1.Start, p2.Start - 1 - p1.Start}, nil
	}
	// code won't reach here
	return nil, nil
}

func (rle *Rles) Add(p RlePair) {
	//perform compaction
	defer rle.compact()

	if len(rle.RlePairs) == 0 || p.Start > rle.RlePairs[len(rle.RlePairs)-1].Start+rle.RlePairs[len(rle.RlePairs)-1].RunLen {
		rle.RlePairs = append(rle.RlePairs, p)
		return
	}
	if p.Start+p.RunLen < rle.RlePairs[0].Start {
		rle.RlePairs = append([]RlePair{p}, rle.RlePairs...)
		return
	}

	_new_rles := make([]RlePair, 0)
	n := len(rle.RlePairs)
	/*
		1. append all values who's elem.end < p.start
		2. check if p has an overlap with last appended value (handle accordingly)
		2a. if it does, perform overlap
		2b. if it doesn't append the value
		3. keep performing check for remaining elem (subSegment/ rSideOverlap)
		4. add remaining elements where elem.start > p.end
	*/

	var i int
	for i = 0; rle.RlePairs[i].Start+rle.RlePairs[i].RunLen < p.Start && i < n; i++ {
		_new_rles = append(_new_rles, rle.RlePairs[i])
	}

	switch {
	case rle.RlePairs[i].isSubSegment(p):
		return
	case rle.RlePairs[i].rSideOverlap(p):
		x := rle.RlePairs[i].overlapReturn(p)
		_new_rles = append(_new_rles, x)
	case rle.RlePairs[i].lSideOverlap(p):
		x := p.overlapReturn(rle.RlePairs[i])
		_new_rles = append(_new_rles, x)
	case p.isSubSegment(rle.RlePairs[i]):
		_new_rles = append(_new_rles, p)
	default:
		_new_rles = append(_new_rles, p)
		_new_rles = append(_new_rles, rle.RlePairs[i:]...)
		rle.RlePairs = _new_rles
		return
	}

	var j int = i
	i++
	for ; i < n; i++ {
		if _new_rles[j].isSubSegment(rle.RlePairs[i]) {
			continue
		} else if _new_rles[j].rSideOverlap(rle.RlePairs[i]) {
			_new_rles[j] = _new_rles[j].overlapReturn(rle.RlePairs[i])
		} else {
			break
		}
	}
	_new_rles = append(_new_rles, rle.RlePairs[i:]...)
	rle.RlePairs = _new_rles
}

func (rle *Rles) Remove(p RlePair) {
	if len(rle.RlePairs) == 0 {
		return
	}

	_new_rles := make([]RlePair, 0)
	n := len(rle.RlePairs)
	/*
		1. iterate through elements where elem.end < p.start
		2. check overlap and split accordingly
		3. keep skipping elements if they are sub segments
		4. check overlap and split accordingly
		5. add remaining elements where elem.start > p.end
	*/
	var i int
	for i = 0; rle.RlePairs[i].Start+rle.RlePairs[i].RunLen < p.Start && i < n; i++ {
		_new_rles = append(_new_rles, rle.RlePairs[i])
	}

	switch {
	case i >= n || rle.RlePairs[i].Start > p.Start+p.RunLen:
		return
	case rle.RlePairs[i].isSubSegment(p):
		before, after := rle.RlePairs[i].splitReturn(p)
		if before != nil {
			_new_rles = append(_new_rles, *before)
		}
		if after != nil {
			_new_rles = append(_new_rles, *after)
		}
		_new_rles = append(_new_rles, rle.RlePairs[i+1:]...)
		rle.RlePairs = _new_rles
		return
	case rle.RlePairs[i].lSideOverlap(p):
		_, toInsert := rle.RlePairs[i].splitReturn(p)
		_new_rles = append(_new_rles, *toInsert)
		_new_rles = append(_new_rles, rle.RlePairs[i+1:]...)
		rle.RlePairs = _new_rles
		return
	case rle.RlePairs[i].rSideOverlap(p):
		toInsert, _ := rle.RlePairs[i].splitReturn(p)
		_new_rles = append(_new_rles, *toInsert)
		i++
	}

	for ; i < n && p.isSubSegment(rle.RlePairs[i]); i++ {
	}

	if i < n {
		if rle.RlePairs[i].lSideOverlap(p) {
			_, toInsert := rle.RlePairs[i].splitReturn(p)
			_new_rles = append(_new_rles, *toInsert)
			_new_rles = append(_new_rles, rle.RlePairs[i+1:]...)
		} else {
			_new_rles = append(_new_rles, rle.RlePairs[i:]...)
		}
	}

	rle.RlePairs = _new_rles
}

//compact causes rle compaction to take place (memory optimzation)
func (rle *Rles) compact() {
	_rleArr := make([]RlePair, 0)
	//iterate through the array and each index checks for overlap with previous one (window search)
	startP := rle.RlePairs[0]
	for i := 1; i < len(rle.RlePairs); i++ {
		if startP.canMerge(rle.RlePairs[i]) {
			startP = startP.mergeReturn(rle.RlePairs[i])
		} else {
			_rleArr = append(_rleArr, startP)
			startP = rle.RlePairs[i]
		}
	}
	_rleArr = append(_rleArr, startP)
	rle.RlePairs = _rleArr
}

func (rle *Rles) Max() (uint16, error) {
	if len(rle.RlePairs) == 0 {
		return 0, fmt.Errorf("EmptyRleError")
	}
	lastPoint := rle.RlePairs[len(rle.RlePairs)-1].Start + rle.RlePairs[len(rle.RlePairs)-1].RunLen
	return lastPoint, nil
}

func (rle *Rles) Min() (uint16, error) {
	if len(rle.RlePairs) == 0 {
		return 0, fmt.Errorf("EmptyRleError")
	}
	firstPoint := rle.RlePairs[0].Start
	return firstPoint, nil
}

func (rle *Rles) NumElem() uint16 {
	currentCount := uint16(0)
	for _, v := range rle.RlePairs {
		currentCount += v.RunLen + 1
	}
	return currentCount
}

func (rle *Rles) Pop() (uint16, error) {
	elem, err := rle.Max()
	if err != nil {
		rle.Remove(RlePair{Start: elem, RunLen: 0})
	}
	return elem, err
}

//Rank returns number of elements -le the given number
func (rle *Rles) Rank(elem uint16) uint16 {
	_total := uint16(0)
	for _, v := range rle.RlePairs {
		if elem > v.Start+v.RunLen {
			_total += v.RunLen + 1
		} else if elem >= v.Start && elem <= v.Start+v.RunLen {
			_total += elem - v.Start + 1
		} else {
			return _total
		}
	}
	return _total
}

//Select returns the element at the i-th index
func (rle *Rles) Select(index uint16) (uint16, error) {
	_indexCount := uint16(0)
	for _, v := range rle.RlePairs {
		if _indexCount+v.RunLen+1 < index {
			_indexCount += v.RunLen + 1
		} else {
			return v.Start + index, nil
		}
	}
	return 0, fmt.Errorf("IndexOutOfBounds")
}

//Index returns the index location of provided element
func (rle *Rles) Index(elem uint16) (int, error) {
	_indexCount := uint16(0)
	for _, v := range rle.RlePairs {
		if v.Start+v.RunLen < elem {
			_indexCount += v.RunLen + 1
		} else if elem >= v.Start && elem <= v.Start+v.RunLen {
			_indexCount += elem - v.Start
			return int(_indexCount), nil
		}
	}
	return 0, fmt.Errorf("ElementNotFound")
}

//TODO -> implement RLE binary operations
func (rle *Rles) Union(rle2 *Rles) Rles {
	_rle := CreateRles()
	var i, j int

	for i < len(rle.RlePairs) && j < len(rle2.RlePairs) {
		if rle.RlePairs[i].Start+rle.RlePairs[i].RunLen < rle2.RlePairs[j].Start {
			_rle.RlePairs = append(_rle.RlePairs, rle.RlePairs[i])
			i++
		} else if rle2.RlePairs[j].Start+rle2.RlePairs[j].RunLen < rle.RlePairs[i].Start {
			_rle.RlePairs = append(_rle.RlePairs, rle2.RlePairs[j])
			j++
		} else {
			if rle.RlePairs[i].isSubSegment(rle2.RlePairs[j]) {
				j++
			} else if rle2.RlePairs[j].isSubSegment(rle.RlePairs[i]) {
				i++
			} else {
				var _overlap RlePair
				if rle.RlePairs[i].lSideOverlap(rle2.RlePairs[j]) {
					_overlap = rle.RlePairs[i].overlapReturn(rle2.RlePairs[j])
				} else {
					_overlap = rle2.RlePairs[j].overlapReturn(rle.RlePairs[i])
				}
				_rle.RlePairs = append(_rle.RlePairs, _overlap)
				i++
				j++
			}
		}
	}

	_rle.RlePairs = append(_rle.RlePairs, rle.RlePairs[i:]...)
	_rle.RlePairs = append(_rle.RlePairs, rle2.RlePairs[j:]...)

	return _rle
}

func (rle *Rles) Intersection(rle2 *Rles) Rles {
	_rle := CreateRles()
	var i, j int

	for i < len(rle.RlePairs) && j < len(rle2.RlePairs) {
		if rle.RlePairs[i].Start+rle.RlePairs[i].RunLen < rle2.RlePairs[j].Start {
			i++
		} else if rle2.RlePairs[j].Start+rle2.RlePairs[j].RunLen < rle.RlePairs[i].Start {
			j++
		} else {
			if rle.RlePairs[i].isSubSegment(rle2.RlePairs[j]) {
				_rle.RlePairs = append(_rle.RlePairs, rle2.RlePairs[j])
				j++
			} else if rle2.RlePairs[j].isSubSegment(rle.RlePairs[i]) {
				_rle.RlePairs = append(_rle.RlePairs, rle.RlePairs[i])
				i++
			} else {
				if rle.RlePairs[i].lSideOverlap(rle2.RlePairs[j]) {
					_rle.RlePairs = append(_rle.RlePairs, rle.RlePairs[i].intersectReturn(rle2.RlePairs[j]))
				} else {
					_rle.RlePairs = append(_rle.RlePairs, rle2.RlePairs[j].intersectReturn(rle.RlePairs[i]))
				}
				i++
				j++
			}
		}
	}

	return _rle
}

func (rle *Rles) Rles2Sarr() Sarr {
	_sarr := CreateSarr()
	var _arr []uint16

	for _, v := range rle.RlePairs {
		for j := v.Start; j <= v.Start+v.RunLen; j++ {
			_arr = append(_arr, j)
		}
	}
	_sarr.Arr = _arr
	return _sarr
}

func (rle *Rles) Rles2Bmps() Bitmaps {
	_bmps := CreateBitmap()

	for _, v := range rle.RlePairs {
		for j := v.Start; j <= v.Start+v.RunLen; j++ {
			_bmps.Add(j)
		}
	}
	return _bmps
}
