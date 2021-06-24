package roar_test

import (
	roar "roar/types"
	"testing"
)

var bitmap_0, bitmap_1 roar.Bitmaps

func TestBitmapSetOps(t *testing.T) {
	bitmap_0 = roar.CreateBitmap()
	bitmap_1 = roar.CreateBitmap()

	bitmap_0.Add(0)
	bitmap_0.Add(1<<16 - 1)
	bitmap_0.Add(31)
	bitmap_0.Add(32)

	bitmap_1.Add(0)
	bitmap_1.Add(15)
	bitmap_1.Add(16)
	bitmap_1.Add(31)
	t.Run("Check Union operation", func(t *testing.T) {
		res := bitmap_0.Union(&bitmap_1)

		required := uint32(0b10000000000000011000000000000001)
		if res.Values[0] != required {
			t.Error("Bitmap Union Failed")
		}
	})
	t.Run("Check Intersection operation", func(t *testing.T) {
		res := bitmap_0.Intersection(&bitmap_1)

		required := uint32(1<<0 | 1<<31)
		if res.Values[0] != required {
			t.Error("Bitmap Intersection Failed")
		}
	})
}
