package sort

import "golang.org/x/exp/constraints"

func merge[T constraints.Ordered](arr []T, left, mid, right int) {
	tep := make([]T, right-left+1)

	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if arr[i] < arr[j] {
			tep[k] = arr[i]
			i++
		} else {
			tep[k] = arr[j]
			j++
		}
		k++
	}
	for i <= mid {
		tep[k] = arr[i]
		i++
		k++
	}
	for j <= right {
		tep[k] = arr[j]
		j++
		k++
	}
	for i := left; i <= right; i++ {
		arr[i] = tep[i-left]
	}
}

func MergeSort[T constraints.Ordered](arr []T) {
	for i := 1; i < len(arr); i *= 2 {
		for j := 0; j < len(arr)-i; j += 2 * i {
			merge(arr, j, j+i-1, min(j+2*i-1, len(arr)-1))
		}
	}
}