package sort

import "golang.org/x/exp/constraints"

func partition[T constraints.Ordered](arr []T, low, high int) int {
	i, j := low, high
	for i < j {
		for arr[i] <= arr[high] && i < high {
			i++
		}
		for arr[j] > arr[high] && j > low {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[high], arr[i] = arr[i], arr[high]
	return i
}

func QuickSort[T constraints.Ordered](arr []T) {
	quicksort(arr, 0, len(arr)-1)
}

func quicksort[T constraints.Ordered](arr []T, low, high int) {
	if low >= high {
		return
	}

	pivot := partition(arr, low, high)
	quicksort(arr, low, pivot-1)
	quicksort(arr, pivot+1, high)
}
