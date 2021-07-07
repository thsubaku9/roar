package roar_test

import (
	roar "roar/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

var rle_0 roar.Rles

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
		assert.Equal(t, uint16(32), res, "Select failed")

		res = rle_0.NumElem()
		assert.Equal(t, uint16(3), res, "NumElem failed")

		res, err = rle_0.Pop()
		assert.Nil(t, err, "Error - Pop failed")
		assert.Equal(t, uint16(32), res, "Select failed")

	})

	t.Run("Check Union operation", func(t *testing.T) {

	})
	t.Run("Check Intersection operation", func(t *testing.T) {

	})
	t.Run("Check SymmetricDifference", func(t *testing.T) {

	})
	t.Run("Check Difference", func(t *testing.T) {

	})

	t.Run("Check Sarr Conversion", func(t *testing.T) {

	})
	t.Run("Check Rles Conversion", func(t *testing.T) {

	})
}
