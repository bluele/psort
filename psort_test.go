package psort

import (
	"fmt"
	"testing"

	"sort"

	"github.com/bluele/randutil"
)

var (
	size = 2000
	k    = 100
)

func makeData(size int) []int {
	vs := make([]int, size)
	for i := 0; i < size; i++ {
		vs[i] = i
	}
	randutil.Shuffle(vs)
	return vs
}

func checkSort(data []int, k int) {
	var prev int
	for i, v := range data[:k] {
		if i > 0 && v < prev {
			panic(fmt.Sprintf("sort error: %v", data[:k]))
		}
		prev = v
	}
}

func BenchmarkSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := makeData(size)
		b.StartTimer()
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
		b.StopTimer()
		checkSort(data, k)
		b.StartTimer()
	}
}

func BenchmarkPSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := makeData(size)
		b.StartTimer()
		ret := SortTopK(data, k)
		b.StopTimer()
		checkSort(ret, k)
		b.StartTimer()
	}
}
