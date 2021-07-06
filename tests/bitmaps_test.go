package roar_test

import (
	roar "roar/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

var bitmap_0, bitmap_1 roar.Bitmaps

func TestBitmapSetOps(t *testing.T) {
	bitmap_0 = roar.CreateBitmap()
	bitmap_1 = roar.CreateBitmap()

	bitmap_0.Add(0)
	bitmap_0.Add(31)
	bitmap_0.Add(32)

	bitmap_1.Add(0)
	bitmap_1.Add(15)
	bitmap_1.Add(16)
	bitmap_1.Add(31)

	t.Run("Unary function tests", func(t *testing.T) {

		res, err := bitmap_0.Max()
		assert.Nil(t, err, "Error triggered in Max")
		assert.Equal(t, uint16(32), res, "Max failed")

		res, err = bitmap_0.Min()
		assert.Nil(t, err, "Error triggered in Min")
		assert.Equal(t, uint16(0), res, "Max failed")

		res = bitmap_0.Rank(32)
		assert.Equal(t, uint16(3), res, "Rank failed")

		res, err = bitmap_0.Select(2)
		assert.Nil(t, err, "Error - Select failed")
		assert.Equal(t, uint16(32), res, "Select failed")

		res = bitmap_0.NumElem()
		assert.Equal(t, uint16(3), res, "NumElem failed")
	})

	t.Run("Check Union operation", func(t *testing.T) {
		res := bitmap_0.Union(&bitmap_1)
		required_0 := uint32(0b10000000000000011000000000000001)
		assert.Equal(t, res.Values[0], required_0, "Bitmap Union Failed")
	})
	t.Run("Check Intersection operation", func(t *testing.T) {
		res := bitmap_0.Intersection(&bitmap_1)
		required_0 := uint32(1<<0 | 1<<31)
		assert.Equal(t, res.Values[0], required_0, "Bitmap Intersection Failed")
	})
	t.Run("Check SymmetricDifference", func(t *testing.T) {
		res := bitmap_0.SymmetricDifference(&bitmap_1)
		required_0 := uint32(1<<15 | 1<<16)
		assert.Equal(t, res.Values[0], required_0, "Bitmap SymmetricDifference Failed")
	})
	t.Run("Check Difference", func(t *testing.T) {
		res := bitmap_0.Difference(&bitmap_1)
		assert.Equal(t, uint32(0), res.Values[0], "Difference Failed")
	})

	t.Run("Check Sarr Conversion", func(t *testing.T) {
		_sarr := bitmap_0.Bmps2Sarr()
		assert.Equal(t, []uint16{0, 31, 32}, _sarr.Arr, "Sarr Conversion failed")
	})
	t.Run("Check Rles Conversion", func(t *testing.T) {
		_rles := bitmap_0.Bmps2Rles()
		_actualRles := roar.CreateRles()
		_actualRles.Add(roar.RlePair{Start: 0x00, RunLen: 0x00})
		_actualRles.Add(roar.RlePair{Start: 0x1f, RunLen: 0x01})
		assert.Equal(t, _actualRles, _rles, "Rles Conversion failed")
	})
}
