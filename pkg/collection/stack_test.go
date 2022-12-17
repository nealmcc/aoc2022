package collection

import (
	"testing"
)

func TestStack_Len(t *testing.T) {
	s := Stack[int]{}

	if s.Len() != 0 {
		t.Log("an empty stack should have length 0")
		t.Fail()
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)

	got := s.Len()
	if got != 3 {
		t.Logf("s.Len() = %d ; want %d", got, 3)
		t.Fail()
	}
}

func TestStack_PushPop(t *testing.T) {
	s := Stack[float64]{}

	s.Push(42.0)

	s.Push(3.0)
	s.Push(2.0)
	s.Push(1.0)

	s.Pop()
	s.Pop()
	s.Pop()

	got, err := s.Pop()
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}

	if got != 42 {
		t.Logf("s.Pop() = %f ; want %f", got, 42.0)
		t.Fail()
	}
}

func TestStack_PopEmpty(t *testing.T) {
	s := Stack[int]{}

	_, err := s.Pop()
	if err == nil {
		t.Log("popping an empty stack should return an error")
		t.Fail()
	}
}

func TestStack_Peek(t *testing.T) {
	type car struct {
		make   string
		model  string
		colour string
	}
	s := Stack[car]{}

	var (
		porsche = car{make: "Porsche", model: "911", colour: "black"}
		tesla   = car{make: "Tesla", model: "Roadster", colour: "red"}
	)

	s.Push(porsche)
	s.Push(tesla)

	s.Pop()

	got, err := s.Peek()
	if err != nil {
		t.Log("unexpected error", err)
		t.Fail()
	}

	if got != porsche {
		t.Logf("s.Peek() = %v ; want %v", got, porsche)
		t.Fail()
	}
}

func TestStack_PeekEmpty(t *testing.T) {
	s := Stack[string]{}

	_, err := s.Peek()
	if err == nil {
		t.Log("peeking at an empty stack should return an error")
		t.Fail()
	}
}
