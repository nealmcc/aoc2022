package collection

// Queue is a generic first-in, first-out container.
// It is not safe to use concurrently.
type Queue[T any] struct {
	data []T
}

// NewQueue initializes a new queue using the given data.
func NewQueue[T any](data ...T) *Queue[T] {
	return &Queue[T]{data: data}
}

// Len returns the number of items in this queue.
func (q *Queue[T]) Len() int {
	return len(q.data)
}

// Push the given item to the back of the queue.
func (q *Queue[T]) Push(x T) {
	(*q).data = append(q.data, x)
}

// Pop returns the front item from the queue, and true iff there is an item.
func (q *Queue[T]) Pop() (T, bool) {
	if len(q.data) == 0 {
		var zero T
		return zero, false
	}

	v := q.data[0]
	q.data = q.data[1:]
	return v, true
}
