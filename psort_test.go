package psort

import (
	"fmt"
	"testing"

	"sort"

	"github.com/bluele/randutil"
)

var (
	size = 20000
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

func checkSorted(data []int, k int) bool {
	var prev int
	for i, v := range data[:k] {
		if i > 0 && v < prev {
			return false
		}
		prev = v
	}
	return true
}

func TestPSort(t *testing.T) {
	for i := 0; i < 100; i++ {
		data := makeData(size)
		Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		}, k)
		if !checkSorted(data, k) {
			t.Fatal(fmt.Sprintf("unsort error: %v", data[:k]))
		}
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
	}
}

func BenchmarkPSort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data := makeData(size)
		b.StartTimer()
		Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		}, k)
	}
}
