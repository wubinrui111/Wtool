package sort

import (
	"math"

	"golang.org/x/exp/constraints"
)

func SelectionSort[T constraints.Ordered](arr []T) {
	n := len(arr)

	for i := range arr {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func Float64SelectionSort(arr []float64) {
	n := len(arr)

	nonNanCount := 0
	for i := 0; i < n; i++ {
		if !math.IsNaN(arr[i]) {
			if i != nonNanCount {
				arr[nonNanCount], arr[i] = arr[i], arr[nonNanCount]
			}
			nonNanCount++
		}
	}

	for i := 0; i < nonNanCount; i++ {
		minIndex := i
		for j := i + 1; j < nonNanCount; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}
