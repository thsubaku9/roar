package roar_test

import (
	roar "roar/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSarrSetOps(t *testing.T) {
	var sarr_0, sarr_1 roar.Sarr

	t.Run("Unary function tests", func(t *testing.T) {
		sarr_0 = roar.CreateSarr()
		sarr_0.Add(0)
		sarr_0.Add(4)
		sarr_0.Add(1)
		sarr_0.Add(2)
		sarr_0.Add(3)

		res, err := sarr_0.Max()
		assert.Nil(t, err, "Error triggered in Max")
		assert.Equal(t, uint16(4), res, "Max failed")

		res, err = sarr_0.Min()
		assert.Nil(t, err, "Error triggered in Min")
		assert.Equal(t, uint16(0), res, "Max failed")

		res = sarr_0.Rank(32)
		assert.Equal(t, uint16(5), res, "Rank failed")

		res, err = sarr_0.Select(2)
		assert.Nil(t, err, "Error - Select failed")
		assert.Equal(t, uint16(2), res, "Select failed")

		res = sarr_0.NumElem()
		assert.Equal(t, uint16(5), res, "NumElem failed")

		res, err = sarr_0.Pop()
		assert.Nil(t, err, "Error - Pop failed")
		assert.Equal(t, uint16(4), res, "Pop failed")
	})

	t.Run("Check Union operation", func(t *testing.T) {
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

		_res := sarr_0.Union(&sarr_1)
		assert.Equal(t, []uint16{0, 1, 2, 3, 4, 5, 6}, _res.Arr, "Sarr Union Failed")
	})

	t.Run("Check Intersection operation", func(t *testing.T) {
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

		_res := sarr_0.Intersection(&sarr_1)
		assert.Equal(t, []uint16{3, 4}, _res.Arr, "Sarr Intersection Failed")
	})

	t.Run("Check Difference operation", func(t *testing.T) {
		sarr_0 = roar.CreateSarr()
		sarr_1 = roar.CreateSarr()
		expected := make([]uint16, 0)
		for i := 0; i < 100; i++ {
			sarr_0.Add(uint16(i))
			if i%2 == 0 {
				sarr_1.Add(uint16(i))
			}
			if i%2 == 1 {
				expected = append(expected, uint16(i))
			}
		}

		_res := sarr_0.Difference(&sarr_1)

		assert.Equal(t, uint16(len(expected)), _res.NumElem(), "Difference Failed")
	})

	t.Run("Check SymmetricDifference operation", func(t *testing.T) {
		sarr_0 = roar.CreateSarr()
		sarr_1 = roar.CreateSarr()
		expected := make([]uint16, 0)
		for i := 0; i < 50; i++ {
			sarr_0.Add(uint16(i))

			if i < 25 {
				expected = append(expected, uint16(i))
			}
		}

		for i := 25; i < 75; i++ {
			sarr_1.Add(uint16(i))
			if !(i < 50) {
				expected = append(expected, uint16(i))
			}
		}

		_res := sarr_0.SymmetricDifference(&sarr_1)
		assert.Equal(t, expected, _res.Arr, "SymmetricDifference Failed")
	})

	t.Run("Check IsDisjoint function", func(t *testing.T) {
		sarr_0 = roar.CreateSarr()
		sarr_1 = roar.CreateSarr()

		for i := 0; i < 10; i++ {
			sarr_0.Add(uint16(i))
			sarr_1.Add(uint16(i + 10))
		}

		_res := sarr_0.IsDisjoint(&sarr_1)
		assert.Equal(t, true, _res, "IsDisjoint Failed for true")

		sarr_1 = roar.CreateSarr()
		for i := 0; i < 3; i++ {
			sarr_1.Add(uint16(i))
		}
		_res = sarr_0.IsDisjoint(&sarr_1)
		assert.Equal(t, false, _res, "IsDisjoint Failed for false")
	})

	t.Run("Check IsSubset function", func(t *testing.T) {
		sarr_0 = roar.CreateSarr()
		sarr_1 = roar.CreateSarr()

		for i := 0; i < 100; i++ {
			sarr_0.Add(uint16(i))
			if i%2 == 0 {
				sarr_1.Add(uint16(i))
			}
		}

		_res := sarr_0.IsSubset(&sarr_1)
		assert.Equal(t, true, _res, "IsSubset Failed for true")

		sarr_1.Add(100)
		_res = sarr_0.IsSubset(&sarr_1)
		assert.Equal(t, false, _res, "IsSubset Failed for false")
	})

	t.Run("Check Clamp function", func(t *testing.T) {
		sarr_0 = roar.CreateSarr()

		sarr_0.Add(1)
		sarr_0.Add(2)
		sarr_0.Add(20)
		sarr_0.Add(80)
		sarr_0.Add(800)

		res := sarr_0.Clamp(10, 100)
		assert.Equal(t, res.NumElem(), uint16(2), "Clamp failed for bitmap")
	})

	t.Run("Check Bitmap conversion", func(t *testing.T) {
		sarr_0 = roar.CreateSarr()
		sarr_0.Add(0)
		sarr_0.Add(4)
		sarr_0.Add(1)
		sarr_0.Add(2)
		sarr_0.Add(3)

		_bmp := sarr_0.Sarr2Bmps()
		_expected := uint32((1 << 0) | (1 << 1) | (1 << 2) | (1 << 3) | (1 << 4))
		assert.Equal(t, _bmp.Values[0], _expected, "Bitmap Conversion Failed")
	})

	t.Run("Check Rle conversion", func(t *testing.T) {
		sarr_0 = roar.CreateSarr()
		sarr_0.Add(0)
		sarr_0.Add(4)
		sarr_0.Add(1)
		sarr_0.Add(2)
		sarr_0.Add(3)

		_rle := sarr_0.Sarr2Rles()
		_expected := []roar.RlePair{{Start: 0, RunLen: 4}}
		assert.Equal(t, _rle.RlePairs, _expected, "Rle Conversion Failed")
	})
}
