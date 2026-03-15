package set

type set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable](items ...T) *set[T] {
	s := &set[T]{
		data: make(map[T]struct{}),
	}

	for _, item := range items {
		s.Add(item)
	}

	return s
}

func (s *set[T]) Add(item T) {
	s.data[item] = struct{}{}
}

func (s *set[T]) Remove(item T) {
	delete(s.data, item)
}

func (s *set[T]) Has(item T) bool {
	_, ok := s.data[item]
	return ok
}

func (s *set[T]) Len() int {
	return len(s.data)
}

func (s *set[T]) Clear() {
	clear(s.data)
}

func (s *set[T]) Items() []T {
	items := make([]T, 0, len(s.data))
	for item := range s.data {
		items = append(items, item)
	}
	return items
}
