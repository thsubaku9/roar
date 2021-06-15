package roar

type Container interface {
	Add(element uint32) error
	Clamp(start, stop uint32)
	Clear()
	Copy() Container
	Debug() string
	Difference() Container
	Remove(element uint32)
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
- Intersection
- isdisjoint
- issubset
- issuperset
- jaccard
- max,min
- numelem
- pop
- rank -> number of elements -le the given number
- remove
- select -> return the element at the i-th index
- symmetric_difference
- union (union_len)
- Overriding default or, and operations


*/
