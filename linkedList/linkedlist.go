package linkedlist

type node[T any] struct {
	last *node[T]
	data T
	next *node[T]
}

type LinkedList[T any] struct {
	head *node[T]
}

func New[T any](items ...T) *LinkedList[T] {
	list := &LinkedList[T]{}
	for _, item := range items {
		list.PushFront(item)
	}
	return list
}

func (l *LinkedList[T]) PushFront(item T) {
	newNode := &node[T]{
		data: item,
	}
	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head.last = newNode
	l.head = newNode
}

func (l *LinkedList[T]) PushBack(item T) {
	newNode := &node[T]{
		data: item,
	}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	newNode.last = current
}

func (l *LinkedList[T]) PopFront() T {
	if l.head == nil {
		panic("Cannot pop from an empty list")
	}

	t := l.head
	result := t.data

	l.head = l.head.next

	if l.head != nil {
		l.head.last = nil
	}

	t.next = nil
	t.last = nil

	return result
}

func (l *LinkedList[T]) GetLast() T {
	if l.head == nil {
		panic("Cannot get last element from an empty list")
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current.data
}

func (l *LinkedList[T]) PopBack() T {
	if l.head == nil {
		panic("Cannot pop from an empty list")
	}

	current := l.head

	if current.next == nil {
		result := current.data
		l.head = nil
		return result
	}

	for current.next != nil {
		current = current.next
	}

	current.last.next = nil

	result := current.data
	current.last = nil

	return result
}

func (l *LinkedList[T]) Len() int {
	count := 0
	current := l.head
	for current != nil {
		count++
		current = current.next
	}
	return count
}

func (l *LinkedList[T]) Empty() bool {
	return l.head == nil
}
