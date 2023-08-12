package eytzinger

import (
	"fmt"
	"math/rand"
	"runtime"
	"slices"
	"testing"
)

var sizes = []int{1, 5, 10, 15, 20, 22, 24, 26, 28, 29, 30}

func makeInts(n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		ints[i] = i
	}
	return ints
}

func BenchmarkEytzinger(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size%d", size), func(b *testing.B) {
			ints := makeInts(1 << size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Eytzinger(ints)
			}
		})
	}
}

func BenchmarkBase(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size%d", size), func(b *testing.B) {
			ints := makeInts(1 << size)
			r := rand.New(rand.NewSource(42))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				target := r.Intn(len(ints))
				runtime.KeepAlive(target)
			}
		})
	}
}

// benchmark for eytzinger.Search
func BenchmarkSearch(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size%d", size), func(b *testing.B) {
			ints := makeInts(1 << size)
			ints = Eytzinger(ints)
			r := rand.New(rand.NewSource(42))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				target := r.Intn(len(ints))
				Search(ints, target)
			}
		})
	}
}

// benchmark for slices.BinarySearch
func BenchmarkStdSearch(b *testing.B) {
	for _, size := range sizes {
		b.Run(fmt.Sprintf("size%d", size), func(b *testing.B) {
			ints := makeInts(1 << size)
			r := rand.New(rand.NewSource(42))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				target := r.Intn(len(ints))
				slices.BinarySearch(ints, target)
			}
		})
	}
}
