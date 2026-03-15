package sort

import (
	"math"
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	// 测试整数切片
	t.Run("Test sorting integers", func(t *testing.T) {
		arr := []int{64, 34, 25, 12, 22, 11, 90}
		expected := []int{11, 12, 22, 25, 34, 64, 90}
		SelectionSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试字符串切片
	t.Run("Test sorting strings", func(t *testing.T) {
		arr := []string{"banana", "apple", "cherry", "date"}
		expected := []string{"apple", "banana", "cherry", "date"}
		SelectionSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试空切片
	t.Run("Test sorting empty slice", func(t *testing.T) {
		arr := []int{}
		expected := []int{}
		SelectionSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试单个元素
	t.Run("Test sorting single element", func(t *testing.T) {
		arr := []int{42}
		expected := []int{42}
		SelectionSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试已排序数组
	t.Run("Test sorting already sorted slice", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		SelectionSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试逆序数组
	t.Run("Test sorting reverse sorted slice", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}
		SelectionSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试浮点数
	t.Run("Test sorting float numbers", func(t *testing.T) {
		arr := []float64{3.5, 1.2, 4.8, 2.1, 0.9}
		expected := []float64{0.9, 1.2, 2.1, 3.5, 4.8}
		SelectionSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试有重复元素的数组
	t.Run("Test sorting slice with duplicates", func(t *testing.T) {
		arr := []int{5, 2, 8, 2, 9, 1, 5}
		expected := []int{1, 2, 2, 5, 5, 8, 9}
		SelectionSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})
}

func TestFloat64SelectionSort(t *testing.T) {
	// 测试包含 NaN 的浮点数数组
	t.Run("Test sorting float64 with NaN", func(t *testing.T) {
		arr := []float64{3.5, 1.2, math.NaN(), 4.8, 2.1, math.NaN(), 0.9}
		Float64SelectionSort(arr)

		// 检查前几个元素是否已排序且不是 NaN
		sortedPart := arr[:len(arr)-2] // 除去两个 NaN 元素
		expectedSorted := []float64{0.9, 1.2, 2.1, 3.5, 4.8}

		if !reflect.DeepEqual(sortedPart, expectedSorted) {
			t.Errorf("Expected %v, got %v", expectedSorted, sortedPart)
		}

		// 检查最后两个元素是否是 NaN
		nanPart := arr[len(arr)-2:]
		for _, v := range nanPart {
			if !math.IsNaN(v) {
				t.Errorf("Expected NaN values at the end, got %v", v)
			}
		}
	})

	// 测试没有 NaN 的浮点数数组
	t.Run("Test sorting float64 without NaN", func(t *testing.T) {
		arr := []float64{3.5, 1.2, 4.8, 2.1, 0.9}
		expected := []float64{0.9, 1.2, 2.1, 3.5, 4.8}
		Float64SelectionSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试全是 NaN 的数组
	t.Run("Test sorting float64 with all NaN", func(t *testing.T) {
		arr := []float64{math.NaN(), math.NaN(), math.NaN()}
		original := make([]float64, len(arr))
		copy(original, arr)
		Float64SelectionSort(arr)

		// 所有元素仍应该是 NaN
		for i, v := range arr {
			if !math.IsNaN(v) || !math.IsNaN(original[i]) {
				t.Errorf("Expected all NaN values to remain NaN after sorting")
			}
		}
	})

	// 测试正负无穷大
	t.Run("Test sorting float64 with infinities", func(t *testing.T) {
		arr := []float64{3.5, math.Inf(1), 1.2, math.Inf(-1), -2.1}
		Float64SelectionSort(arr)
		expected := []float64{math.Inf(-1), -2.1, 1.2, 3.5, math.Inf(1)}
		for i, v := range arr {
			if math.IsInf(expected[i], 0) && math.IsInf(v, 0) &&
				math.Signbit(expected[i]) == math.Signbit(v) {
				continue // 无穷符号相同则跳过
			} else if v != expected[i] {
				t.Errorf("Expected %v, got %v", expected, arr)
				break
			}
		}
	})
}
