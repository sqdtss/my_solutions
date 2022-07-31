package my_template

func QuickSort(sz []int, low, high int) {
	if low < high {
		pivot := partition(sz, low, high)
		QuickSort(sz, low, pivot-1)
		QuickSort(sz, pivot+1, high)
	}
}

func partition(sz []int, low, high int) int {
	pivot := sz[low]
	for low < high {
		for low < high && pivot <= sz[high] {
			high--
		}
		sz[low] = sz[high]
		for low < high && pivot >= sz[low] {
			low++
		}
		sz[high] = sz[low]
	}
	sz[low] = pivot
	return low
}
