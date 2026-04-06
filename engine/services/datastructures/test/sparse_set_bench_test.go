package datastructures_test

import (
	"engine/services/datastructures"
	"testing"
)

const SparseSetMaxIndex uint32 = 16184

// const SparseSetMaxIndex uint32 = 2048

func SparseSetGetBenchmark(b *testing.B, s datastructures.SparseSet[uint32]) {
	for i := range b.N {
		s.Add(uint32(i) % SparseSetMaxIndex)
	}
	b.ResetTimer()
	for i := range b.N {
		s.Get(uint32(i) % SparseSetMaxIndex)
	}
}
func SparseSetGetIndicesBenchmark(b *testing.B, s datastructures.SparseSet[uint32]) {
	for i := range b.N {
		s.Add(uint32(i) % SparseSetMaxIndex)
	}
	b.ResetTimer()

	for b.Loop() {
		s.GetIndices()
	}
}
func SparseSetAddBenchmark(b *testing.B, s datastructures.SparseSet[uint32]) {
	b.ResetTimer()
	for i := 0; i < b.N; {
		for j := uint32(0); j < SparseSetMaxIndex && i < b.N; j++ {
			s.Add(j)
			i++
		}

		b.StopTimer()
		for j := range SparseSetMaxIndex {
			s.Remove(j)
		}
		b.StartTimer()
	}
}
func SparseSetRemoveBenchmark(b *testing.B, s datastructures.SparseSet[uint32]) {
	b.ResetTimer()
	for i := 0; i < b.N; {
		b.StopTimer()
		for j := range SparseSetMaxIndex {
			s.Add(j)
		}
		b.StartTimer()

		for j := uint32(0); j < SparseSetMaxIndex && i < b.N; j++ {
			s.Remove(j)
			i++
		}
	}
}

func BenchmarkSparseSetGetWithoutPaging(b *testing.B) {
	SparseSetGetBenchmark(b, datastructures.NewSparseSet[uint32]())
}
func BenchmarkSparseSetGetWithPaging(b *testing.B) {
	SparseSetGetBenchmark(b, datastructures.NewSparseSetWithPaging[uint32]())
}

func BenchmarkSparseSetGetIndicesWithoutPaging(b *testing.B) {
	SparseSetGetIndicesBenchmark(b, datastructures.NewSparseSet[uint32]())
}
func BenchmarkSparseSetGetIndicesWithPaging(b *testing.B) {
	SparseSetGetIndicesBenchmark(b, datastructures.NewSparseSetWithPaging[uint32]())
}

func BenchmarkSparseSetAddWithoutPaging(b *testing.B) {
	SparseSetAddBenchmark(b, datastructures.NewSparseSet[uint32]())
}
func BenchmarkSparseSetAddWithPaging(b *testing.B) {
	SparseSetAddBenchmark(b, datastructures.NewSparseSetWithPaging[uint32]())
}

func BenchmarkSparseSetRemoveWithoutPaging(b *testing.B) {
	SparseSetRemoveBenchmark(b, datastructures.NewSparseSet[uint32]())
}
func BenchmarkSparseSetRemoveWithPaging(b *testing.B) {
	SparseSetRemoveBenchmark(b, datastructures.NewSparseSetWithPaging[uint32]())
}
