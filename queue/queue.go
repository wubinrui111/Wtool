package queue

type Queue[T any] struct {
	data     []T
	head     int
	tail     int
	count    int
	capacity int
}

func New[T any](items ...T) *Queue[T] {
	if len(items) > 0 {
		q := &Queue[T]{
			data:     make([]T, len(items)*2),
			head:     0,
			tail:     len(items),
			count:    len(items),
			capacity: len(items) * 2,
		}
		copy(q.data, items)
		return q
	}
	initialCapacity := 16
	return &Queue[T]{
		data:     make([]T, initialCapacity),
		head:     0,
		tail:     0,
		count:    0,
		capacity: initialCapacity,
	}
}

func (q *Queue[T]) resize() {
	newData := make([]T, q.capacity*2)

	for i := 0; i < q.count; i++ {
		newData[i] = q.data[(q.head+i)%q.capacity]
	}

	q.data = newData
	q.head = 0
	q.tail = q.count
	q.capacity *= 2
}

func (q *Queue[T]) Enqueue(item T) {
	if q.count == q.capacity {
		q.resize()
	}
	q.data[q.tail] = item
	q.tail = (q.tail + 1) % q.capacity
	q.count++
}

func (q *Queue[T]) Dequeue() T {
	if q.Empty() {
		panic("queue is empty")
	}

	item := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.count--

	return item
}

func (q *Queue[T]) Front() T {
	if q.Empty() {
		panic("queue is empty")
	}
	return q.data[q.head]
}

func (q *Queue[T]) Back() T {
	if q.Empty() {
		panic("queue is empty")
	}
	index := (q.tail - 1 + q.capacity) % q.capacity
	return q.data[index]
}

func (q *Queue[T]) Len() int {
	return q.count
}

func (q *Queue[T]) Empty() bool {
	return q.count == 0
}

func (q *Queue[T]) Clear() {
	q.head = 0
	q.tail = 0
	q.count = 0
	for i := 0; i < q.capacity; i++ {
		var zero T
		q.data[i] = zero
	}
}
