package roar_test

import (
	roar "roar/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sarr_0, sarr_1 roar.Sarr

func InsertElemsIntoSarr(t *testing.T) {
	sarr_0 = roar.CreateSarr()
	sarr_1 = roar.CreateSarr()

	sarr_0.Add(0)
	sarr_0.Add(4)
	sarr_0.Add(1)
	sarr_0.Add(2)
	sarr_0.Add(3)

	sarr_1.Add(3)
	sarr_1.Add(4)
	sarr_1.Add(5)
	sarr_1.Add(6)

	t.Run("Check Union operation", func(t *testing.T) {
		_res := sarr_0.Union(&sarr_1)
		assert.Equal(t, _res, []uint16{0, 1, 2, 3, 4, 5, 6}, "Sarr Union Failed")
	})

	t.Run("Check Intersection operation", func(t *testing.T) {
		_res := sarr_0.Intersection(&sarr_1)
		assert.Equal(t, _res, []uint16{3, 4}, "Sarr Intersection Failed")
	})

	t.Run("Check Bitmap conversion", func(t *testing.T) {
		_bmp := sarr_0.Sarr2Bmps()
		_expected := uint16((1 << 0) | (1 << 1) | (1 << 2) | (1 << 3) | (1 << 4))
		assert.Equal(t, _bmp.Values[0], _expected, "Bitmap Conversion Failed")
	})

	t.Run("Check Rle conversion", func(t *testing.T) {
		_rle := sarr_0.Sarr2Rles()
		_expected := []roar.RlePair{{Start: 0, RunLen: 4}}
		assert.Equal(t, _rle.RlePairs, _expected, "Rle Conversion Failed")
	})
}
