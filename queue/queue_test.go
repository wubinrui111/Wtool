package queue

import (
	"testing"
)

func TestNew(t *testing.T) {
	q := New[int]()
	if q.Len() != 0 {
		t.Errorf("Expected empty queue, got length %d", q.Len())
	}
	if !q.Empty() {
		t.Error("Expected empty queue")
	}

	q2 := New(1, 2, 3)
	if q2.Len() != 3 {
		t.Errorf("Expected queue length 3, got %d", q2.Len())
	}
}

func TestEnqueue(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Len() != 3 {
		t.Errorf("Expected length 3, got %d", q.Len())
	}
	if q.Empty() {
		t.Error("Expected non-empty queue")
	}
}

func TestDequeue(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item := q.Dequeue()
	if item != 1 {
		t.Errorf("Expected to dequeue 1, got %d", item)
	}
	if q.Len() != 2 {
		t.Errorf("Expected length 2 after dequeue, got %d", q.Len())
	}

	item = q.Dequeue()
	if item != 2 {
		t.Errorf("Expected to dequeue 2, got %d", item)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when dequeuing from empty queue")
		}
	}()
	q.Dequeue()
}

func TestFront(t *testing.T) {
	q := New[string]()
	q.Enqueue("first")
	q.Enqueue("second")

	front := q.Front()
	if front != "first" {
		t.Errorf("Expected front to be 'first', got %s", front)
	}

	q.Dequeue()
	front = q.Front()
	if front != "second" {
		t.Errorf("Expected front to be 'second', got %s", front)
	}
}

func TestFrontEmptyQueue(t *testing.T) {
	q := New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when getting front of empty queue")
		}
	}()
	q.Front()
}

func TestBack(t *testing.T) {
	q := New[string]()
	q.Enqueue("first")
	q.Enqueue("second")
	q.Enqueue("third")

	back := q.Back()
	if back != "third" {
		t.Errorf("Expected back to be 'third', got %s", back)
	}

	q.Dequeue() // remove "first"
	back = q.Back()
	if back != "third" {
		t.Errorf("Expected back to still be 'third', got %s", back)
	}
}

func TestBackEmptyQueue(t *testing.T) {
	q := New[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic when getting back of empty queue")
		}
	}()
	q.Back()
}

func TestLenAndEmpty(t *testing.T) {
	q := New[float64]()

	if q.Len() != 0 {
		t.Errorf("Expected length 0, got %d", q.Len())
	}
	if !q.Empty() {
		t.Error("Expected queue to be empty")
	}

	q.Enqueue(1.1)
	if q.Len() != 1 {
		t.Errorf("Expected length 1, got %d", q.Len())
	}
	if q.Empty() {
		t.Error("Expected queue to not be empty")
	}

	q.Dequeue()
	if q.Len() != 0 {
		t.Errorf("Expected length 0 after dequeue, got %d", q.Len())
	}
	if !q.Empty() {
		t.Error("Expected queue to be empty after dequeue")
	}
}

func TestClear(t *testing.T) {
	q := New[int]()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Clear()

	if q.Len() != 0 {
		t.Errorf("Expected length 0 after clear, got %d", q.Len())
	}
	if !q.Empty() {
		t.Error("Expected queue to be empty after clear")
	}
}

func TestResize(t *testing.T) {
	q := New[int]()

	// 填满初始容量直到触发resize
	initialCap := q.capacity
	for i := 0; i < initialCap; i++ {
		q.Enqueue(i)
	}

	// 再添加一个元素，应该触发扩容
	q.Enqueue(initialCap)

	if q.capacity <= initialCap {
		t.Errorf("Expected capacity > %d after resize, got %d", initialCap, q.capacity)
	}

	// 验证所有元素仍然可以正确取出
	for i := 0; i <= initialCap; i++ {
		item := q.Dequeue()
		if item != i {
			t.Errorf("Expected to dequeue %d, got %d", i, item)
		}
	}
}

func TestCircularBuffer(t *testing.T) {
	q := New[int]()

	// 添加一些元素
	for i := 0; i < 5; i++ {
		q.Enqueue(i)
	}

	// 移除一些元素，使头部向前移动
	for i := 0; i < 3; i++ {
		item := q.Dequeue()
		if item != i {
			t.Errorf("Expected to dequeue %d, got %d", i, item)
		}
	}

	// 再添加一些元素，这会使得尾部环绕到数组前面
	for i := 5; i < 10; i++ {
		q.Enqueue(i)
	}

	// 现在队列中的元素应该是 [3, 4, 5, 6, 7, 8, 9]
	expected := []int{3, 4, 5, 6, 7, 8, 9}
	for i, exp := range expected {
		item := q.Dequeue()
		if item != exp {
			t.Errorf("At position %d: expected %d, got %d", i, exp, item)
		}
	}
}
