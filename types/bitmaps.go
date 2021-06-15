package roar

//bitmaps stores 2^16 values in 32bit words -> 1024 entries
type bitmaps struct {
	values []uint32
}

func (bmp *bitmaps) init() {

}
func (bmp *bitmaps) Add(elem uint32) error {
	_ = elem / 32

	return nil
}
