package collection

import (
	"testing"
)

func TestQueue_Len(t *testing.T) {
	q := Queue[int]{}

	if q.Len() != 0 {
		t.Log("an empty queue should have length 0")
		t.Fail()
	}

	q.Push(1)
	q.Push(2)
	q.Push(3)

	got := q.Len()
	if got != 3 {
		t.Logf("s.Len() = %d ; want %d", got, 3)
		t.Fail()
	}
}

func TestQueue_PushPop(t *testing.T) {
	q := Queue[float64]{}

	q.Push(1.0)
	q.Push(2.0)
	q.Push(3.0)
	q.Push(42.0)

	q.Pop()
	q.Pop()
	q.Pop()

	got, ok := q.Pop()
	if got != 42 || !ok {
		t.Logf("s.Pop() = %f, %v ; want %f, %v",
			got, ok, 42.0, true)
		t.Fail()
	}
}

func TestQueue_PopEmpty(t *testing.T) {
	q := Queue[int]{}

	_, ok := q.Pop()
	if ok {
		t.Log("popping an empty queue should return false")
		t.Fail()
	}
}
