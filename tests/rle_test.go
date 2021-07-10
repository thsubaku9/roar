package roar_test

import (
	roar "roar/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rle_0, rle_1 roar.Rles

func TestRleSetOps(t *testing.T) {
	var rle_0 roar.Rles

	t.Run("Unary function tests", func(t *testing.T) {
		rle_0 = roar.CreateRles()

		rle_0.Add(roar.RlePair{Start: 0, RunLen: 1})
		rle_0.Add(roar.RlePair{Start: 1, RunLen: 1})
		rle_0.Add(roar.RlePair{Start: 3, RunLen: 4})
		rle_0.Add(roar.RlePair{Start: 5, RunLen: 2})

		res, err := rle_0.Max()
		assert.Nil(t, err, "Error - Max failed")
		assert.Equal(t, uint16(7), res, "Max failed")

		res, err = rle_0.Min()
		assert.Nil(t, err, "Error - Min failed")
		assert.Equal(t, uint16(0), res, "Max failed")

		res = rle_0.Rank(0)
		assert.Equal(t, uint16(1), res, "Rank failed")
		res = rle_0.Rank(2)
		assert.Equal(t, uint16(3), res, "Rank failed")
		res = rle_0.Rank(32)
		assert.Equal(t, uint16(8), res, "Rank failed")

		res, err = rle_0.Select(2)
		assert.Nil(t, err, "Error - Select failed")
		assert.Equal(t, uint16(2), res, "Select failed")

		res = rle_0.NumElem()
		assert.Equal(t, uint16(8), res, "NumElem failed")

		res, err = rle_0.Pop()
		assert.Nil(t, err, "Error - Pop failed")
		assert.Equal(t, uint16(7), res, "Pop failed")

	})

	t.Run("Check Union operation", func(t *testing.T) {
		rle_0 = roar.CreateRles()
		rle_1 = roar.CreateRles()

		//(10,25),(30,50),(80,100)
		rle_0.Add(roar.RlePair{Start: 10, RunLen: 15})
		rle_0.Add(roar.RlePair{Start: 30, RunLen: 20})
		rle_0.Add(roar.RlePair{Start: 80, RunLen: 20})

		//(5,8),(30,50),(60,70),(75,85),(105,115)
		rle_1.Add(roar.RlePair{Start: 5, RunLen: 3})
		rle_1.Add(roar.RlePair{Start: 30, RunLen: 20})
		rle_1.Add(roar.RlePair{Start: 60, RunLen: 10})
		rle_1.Add(roar.RlePair{Start: 75, RunLen: 10})
		rle_1.Add(roar.RlePair{Start: 105, RunLen: 10})

		res := rle_0.Union(&rle_1)

		expected := roar.CreateRles()

		for _, v := range []roar.RlePair{{5, 3}, {10, 15}, {30, 20}, {60, 10}, {75, 25}, {105, 10}} {
			expected.Add(v)
		}
		assert.Equal(t, expected.RlePairs, res.RlePairs, "Union failed")
	})
	t.Run("Check Intersection operation", func(t *testing.T) {

	})
	t.Run("Check SymmetricDifference", func(t *testing.T) {

	})
	t.Run("Check Difference", func(t *testing.T) {

	})

	t.Run("Check Sarr Conversion", func(t *testing.T) {
		rle_0 = roar.CreateRles()

		rle_0.Add(roar.RlePair{Start: 0, RunLen: 1})
		rle_0.Add(roar.RlePair{Start: 3, RunLen: 4})
		rle_0.Add(roar.RlePair{Start: 10, RunLen: 2})

		_sarr := rle_0.Rles2Sarr()

		assert.Equal(t, []uint16{0, 1, 3, 4, 5, 6, 7, 10, 11, 12}, _sarr.Arr, "Sarr conversion failed")
	})
	t.Run("Check Rles Conversion", func(t *testing.T) {
		rle_0 = roar.CreateRles()

		rle_0.Add(roar.RlePair{Start: 0, RunLen: 1})
		rle_0.Add(roar.RlePair{Start: 3, RunLen: 4})
		rle_0.Add(roar.RlePair{Start: 10, RunLen: 2})

		_bmps := rle_0.Rles2Bmps()
		_sarr := rle_0.Rles2Sarr()
		_reqBmps := _sarr.Sarr2Bmps()
		assert.Equal(t, _reqBmps.Values[0], _bmps.Values[0], "Bmps conversion failed")
	})
}
