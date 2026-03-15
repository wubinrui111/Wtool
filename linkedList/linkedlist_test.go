package linkedlist

import "testing"

func TestNew(t *testing.T) {
	list := New[int]()
	if list.head != nil {
		t.Error("Expected empty list")
	}

	listWithItems := New[int](1, 2, 3)
	if listWithItems.head == nil {
		t.Error("Expected non-empty list")
	}
}

func TestPushFront(t *testing.T) {
	list := New[int]()
	list.PushFront(1)
	list.PushFront(2)
	list.PushFront(3)

	if list.head.data != 3 {
		t.Errorf("Expected head to be 3, got %d", list.head.data)
	}
	if list.head.next.data != 2 {
		t.Errorf("Expected second element to be 2, got %d", list.head.next.data)
	}
	if list.head.next.next.data != 1 {
		t.Errorf("Expected third element to be 1, got %d", list.head.next.next.data)
	}
}

func TestPushBack(t *testing.T) {
	list := New[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	if list.head.data != 1 {
		t.Errorf("Expected head to be 1, got %d", list.head.data)
	}
	if list.head.next.data != 2 {
		t.Errorf("Expected second element to be 2, got %d", list.head.next.data)
	}
	if list.head.next.next.data != 3 {
		t.Errorf("Expected third element to be 3, got %d", list.head.next.next.data)
	}
}

func TestPopFront(t *testing.T) {
	list := New[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	item := list.PopFront()
	if item != 1 {
		t.Errorf("Expected popped item to be 1, got %d", item)
	}

	item = list.PopFront()
	if item != 2 {
		t.Errorf("Expected popped item to be 2, got %d", item)
	}

	item = list.PopFront()
	if item != 3 {
		t.Errorf("Expected popped item to be 3, got %d", item)
	}
}

func TestGetLast(t *testing.T) {
	list := New[int]()
	list.PushFront(1)
	list.PushBack(2)
	list.PushBack(3)

	last := list.GetLast()
	if last != 3 {
		t.Errorf("Expected last element to be 3, got %d", last)
	}

	list.PushBack(4)
	last = list.GetLast()
	if last != 4 {
		t.Errorf("Expected last element to be 4, got %d", last)
	}
}

func TestPopBack(t *testing.T) {
	list := New[int]()
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	item := list.PopBack()
	if item != 3 {
		t.Errorf("Expected popped item to be 3, got %d", item)
	}

	item = list.PopBack()
	if item != 2 {
		t.Errorf("Expected popped item to be 2, got %d", item)
	}

	item = list.PopBack()
	if item != 1 {
		t.Errorf("Expected popped item to be 1, got %d", item)
	}
}

func TestMixedOperations(t *testing.T) {
	list := New[int]()
	list.PushFront(1)
	list.PushBack(2)
	list.PushFront(3)
	list.PushBack(4)

	expected := []int{3, 1, 2, 4}
	current := list.head
	i := 0
	for current != nil {
		if current.data != expected[i] {
			t.Errorf("At index %d: expected %d, got %d", i, expected[i], current.data)
		}
		current = current.next
		i++
	}

	if i != len(expected) {
		t.Errorf("Expected %d elements, got %d", len(expected), i)
	}
}

func TestSingleElement(t *testing.T) {
	list := New[int]()
	list.PushBack(42)

	if list.head.data != 42 {
		t.Errorf("Expected single element to be 42, got %d", list.head.data)
	}
	if list.head.next != nil {
		t.Error("Expected single element list to have no next node")
	}
	if list.head.last != nil {
		t.Error("Expected single element list to have no last node")
	}
}