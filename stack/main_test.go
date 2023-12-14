package main

import "testing"

func TestCreateStack(t *testing.T) {
	s := NewStack[int](1)
	t.Log(s.Capacity(), s.Size())
}

func TestPush(t *testing.T) {
	s := NewStack[int](4)
	t.Log(s.Capacity(), s.Size())
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	t.Log(s.Capacity(), s.Size())
	d, e := s.Pop()
	t.Log(e, d)
	d, e = s.Pop()
	t.Log(e, d)
	d, e = s.Pop()
	t.Log(e, d)
	d, e = s.Pop()
	t.Log(e, d)
	t.Log(s.Capacity(), s.Size())
	d, e = s.Pop()
	t.Log(e, d)
	t.Log(s.Capacity(), s.Size())
}
