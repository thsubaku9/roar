package roar

//Container acts as the top level object from which all interactions are done
type Container interface {
	Add(element uint32) error
	Clamp(start, stop uint32)
	Clear()
	Copy() Container
	Debug() string
	Difference() Container
	Intersection(con Container) (Container, error)
	Pop() (uint32, error) //removes the element with highest value
	Remove(element uint32)
	Select(index uint32) (uint32, error) //return the element at the i-th index
	Symmetric_difference(con Container) (Container, error)
	Union(con Container) (Container, error)
}

type RoaringBitmap struct {
	key   []uint16
	value []Container
}

//Roar returns a new RoaringBitmap
func Roar(values ...uint32) RoaringBitmap {
	return RoaringBitmap{key: nil, value: nil}
}

/*

- Flip_range
- Index
- isdisjoint
- issubset
- issuperset
- jaccard
- max,min
- numelem
- pop
- rank -> number of elements -le the given number
- remove
- Overriding default or, and operations


*/
