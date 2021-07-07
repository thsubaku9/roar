package roar

import (
	. "roar/util"
)

//Container acts as the top level object from which all interactions are done
type Container interface {
	Add(element uint32)
	Clamp(start, stop uint32)
	Clear()
	Copy() Container
	Debug() string
	Difference() Container
	FlipRange(start, stop uint32)
	Index(element uint32) (int, error) //returns the index location of provided element
	Intersection(con Container) (Container, error)
	IsDisjoint(con Container) bool
	IsSubset(con Container) bool
	IsSuperset(con Container) bool
	Jaccard(con Container) float32
	Max() (uint32, error)
	Min() (uint32, error)
	NumElem() uint32
	Pop() (uint32, error)       //removes the element with highest value
	Rank(element uint32) uint32 //number of elements -le the given number
	Remove(element uint32)
	Select(index uint32) (uint32, error) //return the element at the i-th index
	SymmetricDifference(con Container) Container
	Union(con Container) (Container, error)
}

type SubContainer interface {
	Add(element uint16)
	Difference(sub SubContainer) SubContainer
	Index(element uint16) (int, error) //returns the index location of provided element
	Intersection(sub SubContainer) (SubContainer, error)
	IsDisjoint(sub SubContainer) bool
	IsSubset(sub SubContainer) bool
	IsSuperset(sub SubContainer) bool
	Max() (uint16, error)
	Min() (uint16, error)
	NumElem() uint16
	Pop() (uint16, error)       //removes the element with highest value
	Rank(element uint16) uint16 //number of elements -le the given number
	Remove(element uint16)
	SymmetricDifference(sub SubContainer) SubContainer
	Union(sub SubContainer) (SubContainer, error)
}

//TODO - sub container conversion will depend on current size of sub container vs alternatives
type RoaringBitmap struct {
	subContainers []*SubContainer
	numElem       uint32
}

//Roar returns a new RoaringBitmap
func Roar(values ...uint32) RoaringBitmap {
	return RoaringBitmap{subContainers: make([]*SubContainer, 16), numElem: 0}
}

func (r *RoaringBitmap) Add(element uint32) {
	key, val := int(element/SplitVal), uint16(element%SplitVal)
	if r.subContainers[key] == nil {
		// set a default container type
		r.subContainers[key] = new(SubContainer)
	}
	(*r.subContainers[key]).Add(val)
}

func (r *RoaringBitmap) Remove(element uint32) {
	key, val := int(element/SplitVal), uint16(element%SplitVal)
	if r.subContainers[key] == nil {
		// set a default container type
		r.subContainers[key] = new(SubContainer)
	}
	(*r.subContainers[key]).Remove(val)
}

/* TODO -> figure out inter sub container conversion criteria.

Sarr -> Bitmaps if (2^16)/16 elements have been inserted
Sarr -> RLE is 2^16 * 16 vs  2^16 * 32 (criteria needs to be fleshed out)
*/

//TODO -> need to add another structure for abstraction since interface method params need to be of type interface and not struct
