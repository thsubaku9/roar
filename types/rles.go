package roar

import (
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
	return (p2.Start+p2.RunLen >= p1.Start) && p2.Start+p2.RunLen <= p1.Start+p1.RunLen
}

func (p1 RlePair) rSideOverlap(p2 RlePair) bool {
	return (p2.Start <= p1.Start+p1.RunLen) && (p2.Start >= p1.Start)
}

func (p1 RlePair) isSubSegment(p2 RlePair) bool {
	return p1.Start <= p2.Start && p1.Start+p1.RunLen >= p2.Start+p2.RunLen
}

//mergeReturn assumes the two pairs do overlap and combines them
func (p1 RlePair) overlapReturn(p2 RlePair) RlePair {
	return RlePair{p1.Start, p1.RunLen + p2.RunLen - (p1.Start + p1.RunLen - p2.Start)}
}

//canMerge checks not overlap, but successive sequences for given pairs
func (p1 RlePair) canMerge(p2 RlePair) bool {
	return p2.Start-p1.Start+p1.RunLen == 1 || p1.Start-p2.Start+p2.RunLen == 1
}

func (p1 RlePair) splitReturn(p2 RlePair) []RlePair {
	//TODO -> implement splitReturn for Remove func
	return []RlePair{p1}
}

func (rle *Rles) Add(p RlePair) {
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
	*/

	var i int
	for i = 0; rle.RlePairs[i].Start+rle.RlePairs[i].RunLen < p.Start && i < n; i++ {
		_new_rles = append(_new_rles, rle.RlePairs[i])
	}

	if rle.RlePairs[i].isSubSegment(p) {
		return
	} else if rle.RlePairs[i].rSideOverlap(p) {
		x := rle.RlePairs[i].overlapReturn(p)
		_new_rles = append(_new_rles, x)
	} else if rle.RlePairs[i].lSideOverlap(p) {
		x := p.overlapReturn(rle.RlePairs[i])
		_new_rles = append(_new_rles, x)
	} else if p.isSubSegment(rle.RlePairs[i]) {
		_new_rles = append(_new_rles, p)
	} else {
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
			_new_rles = append(_new_rles, rle.RlePairs[i])
			j++
		}
	}
	rle.RlePairs = _new_rles
}

func (rle *Rles) Remove(p RlePair) {
	if len(rle.RlePairs) == 0 {
		return
	}
	/*
		1. iterate through elements where elem.end < p.start
		2. ???
	*/

	//_new_rles := make([]RlePair, 0)
}

func (rle *Rles) Union(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}

func (rle *Rles) Intersection(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}
