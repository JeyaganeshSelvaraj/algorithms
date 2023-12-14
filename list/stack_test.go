package main

import "testing"

func TestStackPush(t *testing.T) {
	s, err := NewStack[int](10)
	if err != nil {
		t.FailNow()
	}
	if err := s.Push(1); err != nil {
		t.FailNow()
	} else {
		t.Log(s.String())
	}
	s.Push(2)
	s.Push(3)
	s.Push(4)
	t.Log(s.String())
}
func TestPop(t *testing.T) {
	s, err := NewStack[int](10)
	if err != nil {
		t.FailNow()
	}
	if _, err := s.Pop(); err == nil {
		t.FailNow()
	}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	if data, err := s.Pop(); err != nil {
		t.FailNow()
	} else {
		t.Log(data)
	}
	s.Pop()
	s.Pop()
	t.Log(s.String())
	s.Pop()
	t.Log(s.String())
	if _, err := s.Pop(); err == nil {
		t.FailNow()
	}
}
