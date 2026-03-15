package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	// 测试整数切片
	t.Run("Test sorting integers", func(t *testing.T) {
		arr := []int{64, 34, 25, 12, 22, 11, 90}
		expected := []int{11, 12, 22, 25, 34, 64, 90}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试字符串切片
	t.Run("Test sorting strings", func(t *testing.T) {
		arr := []string{"banana", "apple", "cherry", "date"}
		expected := []string{"apple", "banana", "cherry", "date"}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试空切片
	t.Run("Test sorting empty slice", func(t *testing.T) {
		arr := []int{}
		expected := []int{}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试单个元素
	t.Run("Test sorting single element", func(t *testing.T) {
		arr := []int{42}
		expected := []int{42}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试已排序数组
	t.Run("Test sorting already sorted slice", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		expected := []int{1, 2, 3, 4, 5}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试逆序数组
	t.Run("Test sorting reverse sorted slice", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试浮点数
	t.Run("Test sorting float numbers", func(t *testing.T) {
		arr := []float64{3.5, 1.2, 4.8, 2.1, 0.9}
		expected := []float64{0.9, 1.2, 2.1, 3.5, 4.8}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试有重复元素的数组
	t.Run("Test sorting slice with duplicates", func(t *testing.T) {
		arr := []int{5, 2, 8, 2, 9, 1, 5}
		expected := []int{1, 2, 2, 5, 5, 8, 9}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试较大的数组
	t.Run("Test sorting larger array", func(t *testing.T) {
		arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
		expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})

	// 测试字符切片
	t.Run("Test sorting byte slice", func(t *testing.T) {
		arr := []byte{'d', 'a', 'c', 'b'}
		expected := []byte{'a', 'b', 'c', 'd'}
		QuickSort(arr)
		if !reflect.DeepEqual(arr, expected) {
			t.Errorf("Expected %v, got %v", expected, arr)
		}
	})
}