package main

import (
	"errors"
	"fmt"
	"strings"
)

type Stack[T any] struct {
	head     ListElement[T]
	len      int
	capacity int
}

func (s *Stack[T]) Len() int {
	return s.len
}

func (s *Stack[T]) Capacity() int {
	return s.capacity
}
func NewStack[T any](capacity int) (*Stack[T], error) {
	if capacity <= 0 {
		return nil, errors.New("capacity must be greater than 0")
	}
	return &Stack[T]{
		capacity: capacity,
	}, nil
}
func (s *Stack[T]) Push(data T) error {
	if s.len == s.capacity {
		return errors.New("stack is full")
	}
	elem := &ListElement[T]{Data: data}
	if s.len != 0 {
		h := s.head
		elem.Next = &h
	}
	s.head = *elem
	s.len++
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	var res T
	if s.len == 0 {
		return res, errors.New("stack is empty")
	}
	head := s.head
	res = s.head.Data
	//last data in the stack
	if head.Next != nil {
		s.head = *head.Next
	} else {
		s.head = ListElement[T]{}
	}
	s.len--
	return res, nil
}
func (s *Stack[T]) String() string {
	str := strings.Builder{}
	node := &s.head
	for {
		str.WriteString(fmt.Sprintf("%v -> ", node.Data))
		if node.Next == nil {
			str.WriteString("nil")
			break
		}
		node = node.Next
	}
	return str.String()
}
