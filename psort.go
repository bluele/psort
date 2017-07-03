package psort

import (
	"reflect"
)

type (
	swapFunc func(int, int)
	compare  func(i, j int) bool
)

// Slice sorts top k items of the provided slice given the provided less function.
//
// The function panics if the provided interface is not a slice.
func Slice(slice interface{}, less compare, k int) {
	var index, rank, start int

	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	end := rv.Len() - 1

	for end > start {
		index = partition(&rv, swap, less, start, end)
		rank = index + 1
		if rank >= k {
			end = index - 1
		} else if (index - start) > (end - index) {
			quickSort(&rv, swap, less, index+1, end)
			end = index - 1
		} else {
			quickSort(&rv, swap, less, start, index-1)
			start = index + 1
		}
	}
}

func partition(rv *reflect.Value, swap swapFunc, less compare, start, end int) int {
	i := start
	for j := start + 1; j <= end; j++ {
		if less(j, start) {
			i++
			swap(i, j)
		}
	}
	swap(i, start)
	return i
}

func quickSort(rv *reflect.Value, swap swapFunc, less compare, start, end int) {
	if start < end {
		index := partition(rv, swap, less, start, end)
		quickSort(rv, swap, less, start, index-1)
		quickSort(rv, swap, less, index+1, end)
	}
}
