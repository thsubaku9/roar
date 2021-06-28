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
	return (p2.Start <= p1.Start+p1.RunLen) && (p1.Start <= p2.Start)
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
	if len(rle.RlePairs) == 0 {
		rle.RlePairs = append(rle.RlePairs, p)
		return
	}
	_new_rles := make([]RlePair, 0)

	//starting insertion
	if p.Start < rle.RlePairs[0].Start {
		_new_rles = append(_new_rles, p)

		var i int
		for i, _ = range rle.RlePairs {
			//check if it is a subsegment
			if p.isSubSegment(rle.RlePairs[i]) {
				continue
			}
			//check for roverlap

			// if neither then append directly to _new_rles
		}
		rle.RlePairs = _new_rles
	} else if p.Start > rle.RlePairs[len(rle.RlePairs)-1].Start+rle.RlePairs[len(rle.RlePairs)-1].RunLen {
		rle.RlePairs = append(rle.RlePairs, p)
	} else {
		var i int
		for i, _ = range rle.RlePairs {
			if rle.RlePairs[i].Start > p.Start {
				break
			}
		}
		//this is the insertion point, check for loverlap. after inserting keep checking for roverlap/subsegment until there isn't any
	}
	//iterate through all elems, in case of overlap combine properly
}

func (rle *Rles) Remove(p RlePair) {
	if len(rle.RlePairs) == 0 {
		return
	}
	//iterate through all elems, in case of overlap split properly
}

func (rle *Rles) Union(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}

func (rle *Rles) Intersection(rle2 *Rles) Rles {
	_rle := CreateRles()

	return _rle
}
