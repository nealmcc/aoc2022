package collection

import (
	"errors"
	"sync"
)

// Stack is a generic first-in, last-out container.
// It is safe to use concurrently.
type Stack[T any] struct {
	mu   sync.RWMutex
	data []T
}

// Len returns the number of items in this stack.
func (s *Stack[T]) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

// Push the given item to the top of the stack.
func (s *Stack[T]) Push(x T) {
	s.mu.Lock()
	s.data = append(s.data, x)
	s.mu.Unlock()
}

// Pop returns the top item from the stack, or an error if the stack is empty.
func (s *Stack[T]) Pop() (T, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.data) == 0 {
		var zero T
		return zero, errors.New("cannot pop an empty stack")
	}

	last := len(s.data) - 1
	v := s.data[last]
	s.data = s.data[:last]
	return v, nil
}

// Peek returns the top item from the stack, or an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.data) == 0 {
		var zero T
		return zero, errors.New("cannot peek at an empty stack")
	}

	return s.data[len(s.data)-1], nil
}
