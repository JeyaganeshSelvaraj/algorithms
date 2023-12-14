package main

import "errors"

type Stack[T any] struct {
	capacity int
	elements []T
}

func NewStack[T any](capacity int) *Stack[T] {
	return &Stack[T]{
		capacity: capacity,
	}
}

func (s *Stack[T]) Capacity() int {
	return s.capacity
}
func (s *Stack[T]) Size() int {
	return len(s.elements)
}
func (s *Stack[T]) Push(val T) error {
	if len(s.elements) == s.capacity {
		return errors.New("Stack is full")
	}
	s.elements = append(s.elements, val)
	return nil
}
func (s *Stack[T]) Pop() (T, error) {
	var res T
	size := len(s.elements)
	if size == 0 {
		return res, errors.New("Stack is full")
	}
	res = s.elements[size-1]
	s.elements = s.elements[:size-1]
	return res, nil
}
