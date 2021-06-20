package roar

//Container acts as the top level object from which all interactions are done
type Container interface {
	Add(element uint32) error
	Clamp(start, stop uint32)
	Clear()
	Copy() Container
	Debug() string
	Difference() Container
	FlipRange(start, stop uint32)
	Index(element uint32) (uint32, error) //returns the index location of provided element
	Intersection(con Container) (Container, error)
	IsDisjoint(con Container) bool
	IsSubset(con Container) bool
	IsSuperset(con Container) bool
	Jaccard(con Container) float32
	Max() (uint32, error)
	Min() (uint32, error)
	Pop() (uint32, error)         //removes the element with highest value
	Rank(element uint32) []uint32 //number of elements -le the given number
	Remove(element uint32) error
	Select(index uint32) (uint32, error) //return the element at the i-th index
	SymmetricDifference(con Container) (Container, error)
	Union(con Container) (Container, error)
}

//TODO - Overriding default or, and operations

//TODO - sub container conversion will depend on current size of sub container vs alternatives
type RoaringBitmap struct {
	key     []uint16
	value   []Container
	numElem uint32
}

//Roar returns a new RoaringBitmap
func Roar(values ...uint32) RoaringBitmap {
	return RoaringBitmap{key: nil, value: nil, numElem: 0}
}
