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
	var minP, maxP RlePair = p1, p2
	if p1.Start > p2.Start {
		minP = p2
		maxP = p1
	}
	return RlePair{minP.Start, minP.RunLen + maxP.RunLen - (minP.Start + minP.RunLen - maxP.Start)}
}

//canMerge checks not overlap, but successive sequences for given pairs
func (p1 RlePair) canMerge(p2 RlePair) bool {
	return p2.Start-p1.Start+p1.RunLen == 1 || p1.Start-p2.Start+p2.RunLen == 1
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
		_new_rles = append(_new_rles, *before, *after)
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

func (rle *Rles) Union(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}

func (rle *Rles) Intersection(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}
