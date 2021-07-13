package roar_bench

import (
	roar "roar/types"
	"runtime"
	"testing"
)

//memory usage bench
// go test -bench  BenchmarkMemConsumption
func BenchmarkMemConsumption(b *testing.B) {
	b.StopTimer()
	sarr_0 := roar.CreateSarr()
	N := uint16(1e3)
	for i := uint16(0); i < N; i++ {
		sarr_0.Add(i)
	}
	for i2 := 2 * N; i2 >= N; i2-- {
		sarr_0.Add(i2)
	}
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	b.Logf("HeapInUse: %d, HeapObjects: %d", stats.HeapInuse, stats.HeapObjects)
	b.StartTimer()
}

//speed bench
//binary operations
//merging 2 large Roars
func BenchmarkUnionSarr(b *testing.B) {
	b.StopTimer()
	sarr_0 := roar.CreateSarr()
	N := uint16(1e3)
	for i := uint16(0); i < N; i++ {
		sarr_0.Add(i)
	}

	sarr_1 := roar.CreateSarr()
	for i := N; i < 2*N; i++ {
		sarr_1.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sarr_0.Union(&sarr_1)
	}
}

func BenchmarkIntersectSarr(b *testing.B) {
	b.StopTimer()
	sarr_0 := roar.CreateSarr()
	N := uint16(1e3)
	for i := uint16(0); i < N; i++ {
		sarr_0.Add(i)
	}

	sarr_1 := roar.CreateSarr()
	for i := uint16(0); i < 2*N; i += 2 {
		sarr_1.Add(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		sarr_0.Union(&sarr_1)
	}
}
