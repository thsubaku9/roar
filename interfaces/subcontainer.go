package roar

import (
	"roar/util"
)

type SubContainer interface {
	Add(element uint16)
	Clamp(start, stop uint16) SubContainer
	Difference(sub SubContainer) SubContainer
	Index(element uint16) (int, error) //returns the index location of provided element
	// Intersection(sub SubContainer) (SubContainer, error)
	IsDisjoint(sub SubContainer) bool
	IsSubset(sub SubContainer) bool
	IsSuperset(sub SubContainer) bool
	Max() (uint16, error)
	Min() (uint16, error)
	NumElem() uint16
	Pop() (uint16, error)       //removes the element with highest value
	Rank(element uint16) uint16 //number of elements -le the given number
	Remove(element uint16)
	// SymmetricDifference(sub SubContainer) SubContainer
	// Union(sub SubContainer) (SubContainer, error)
	ScType() util.SubContainerType
}
