package stack_test

import (
	"testing"
	"wtool/stack"
)

func TestStack(t *testing.T) {
	stack := stack.New[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.Push(7)
	stack.Push(8)
	stack.Push(9)
	stack.Push(10)
	if stack.Len() != 10 {
		t.Errorf("stack length is not 10")
	}
	for i := 10; i > 0; i-- {
		if stack.Len() != i {
			t.Errorf("stack length is not %d", i)
		}
		if stack.Pop() != i {
			t.Errorf("stack pop is not %d", i)
		}
	}
}

func TestStackAllNew(t *testing.T) {
	stack := stack.New(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	if stack.Len() != 10 {
		t.Errorf("stack length is not 10")
	}
	for i := 10; i > 0; i-- {
		if stack.Len() != i {
			t.Errorf("stack length is not %d", i)
		}
		if stack.Pop() != i {
			t.Errorf("stack pop is not %d", i)
		}
	}
}
