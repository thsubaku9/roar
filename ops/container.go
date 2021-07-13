package roar

// import (
// 	. "roar/util"
// )

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

//TODO - sub container conversion will depend on current size of sub container vs alternatives
type container struct {
	subContainers []*SubContainer
	numElem       uint32
}

/*
//Roar returns a new RoaringBitmap (Container)
func Roar(values ...uint32) Container {
	res := container{subContainers: make([]*SubContainer, 16), numElem: 0}

	return &res
}

func (r *container) Add(element uint32) {
	key, val := int(element/SplitVal), uint16(element%SplitVal)
	if r.subContainers[key] == nil {
		// set a default container type
		r.subContainers[key] = new(SubContainer)
	}
	(*r.subContainers[key]).Add(val)
}

func (r *container) Clamp(start, stop uint32) {

}

func (r *container) Clear() {
	r = &container{subContainers: make([]*SubContainer, 16), numElem: 0}
}

func (r *container) Debug() string {
	return fmt.Sprintf("Containers - %v, NumElem - %v", r.subContainers, r.numElem)
}

func (r *container) Copy() Container {
	_r := &container{subContainers: make([]*SubContainer, 16), numElem: r.numElem}

	for i, v := range r.subContainers {
		if v != nil {
			//TODO -> exact copy based on underlying implementation
		}
	}

	return _r
}

func (r *container) Remove(element uint32) {
	key, val := int(element/SplitVal), uint16(element%SplitVal)
	if r.subContainers[key] == nil {
		// set a default container type
		r.subContainers[key] = new(SubContainer)
	}
	(*r.subContainers[key]).Remove(val)
}

//TODO -> need to add another structure for abstraction since interface method params need to be of type interface and not struct

*/
