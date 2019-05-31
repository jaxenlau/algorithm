package stack

import (
	"testing"
)

func Test(t *testing.T) {
	t.Log("start")
	s := New(10)

	if s.IsFull() {
		t.Error("New Stack length should be 0")
	}

	s.Push(1)

	if (*s.Top()).(int) != 1 {
		t.Error("top element should be 1")
	}

	if (*s.Pop()).(int) != 1 {
		t.Error("Pop the top element should be 1")
	}

	if !s.IsEmpty() {
		t.Error("Stack length should be 0")
	}

	s.Push(2)
	s.Push(5)

	if (*s.Pop()).(int) != 5 {
		t.Error("Pop the top element should be 1")
	}

	if (*s.Pop()).(int) != 2 {
		t.Error("Pop the top element should be 1")
	}

	if !s.IsEmpty() {
		t.Error("Stack length should be 0")
	}
	t.Log("end")
}
