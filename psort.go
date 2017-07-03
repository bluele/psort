package psort

// SortTopK returns sorted top k items.
func SortTopK(data []int, k int) []int {
	var index, rank, start int
	end := len(data) - 1

	for end > start {
		index = partition(data, start, end)
		rank = index + 1
		if rank >= k {
			end = index - 1
		} else if (index - start) > (end - index) {
			quickSort(data, index+1, end)
			end = index - 1
		} else {
			quickSort(data, start, index-1)
			start = index + 1
		}
	}
	return data[:k]
}

func partition(data []int, start, end int) int {
	x := data[start]
	i := start
	for j := start + 1; j <= end; j++ {
		if x > data[j] {
			i++
			data[i], data[j] = data[j], data[i]
		}
	}
	data[start], data[i] = data[i], data[start]
	return i
}

func quickSort(data []int, start, end int) {
	if start < end {
		index := partition(data, start, end)
		quickSort(data, start, index-1)
		quickSort(data, index+1, end)
	}
}
