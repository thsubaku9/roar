package roar

type Iterator interface {
	Next() interface{}
	HasNext() bool
}
