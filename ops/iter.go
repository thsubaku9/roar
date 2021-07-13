package roar

type Iterator interface {
	next() interface{}
	hasNext() bool
}
