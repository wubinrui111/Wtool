package stack

type Stack[T any] struct {
	data []T
	top  int
}

func New[T any](n ...T) *Stack[T] {
	if n != nil {
		return &Stack[T]{
			data: n,
			top:  len(n),
		}
	}
	return &Stack[T]{
		data: make([]T, 0, 16),
		top:  0,
	}
}

func (s *Stack[T]) Pop() T {
	if s.top == 0 {
		panic("stack is empty")
	}
	t := s.data[s.top-1]
	s.data = s.data[:s.top]
	s.top--
	return t
}

func (s *Stack[T]) Push(v T) {
	s.data = append(s.data, v)
	s.top++
}

func (s *Stack[T]) Peek() T {
	if s.top == 0 {
		panic("stack is empty")
	}
	return s.data[s.top-1]
}

func (s *Stack[T]) Len() int {
	return s.top
}

func (s *Stack[T]) Empty() bool {
	return s.top == 0
}
