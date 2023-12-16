package main

import (
	"errors"
	"fmt"
	"strings"
)

type Element[T any] struct {
	Data  T
	Prev  *Element[T]
	Next  *Element[T]
	Child *Element[T]
}
type DoubleLinkedList[T any] struct {
	root Element[T]
}

func CreateDoubleLinkedList[T any](vals []T) *DoubleLinkedList[T] {
	lst := &DoubleLinkedList[T]{}
	if len(vals) == 0 {
		return lst
	}
	lst.root = Element[T]{
		Data: vals[0],
	}
	curr := &lst.root
	for _, val := range vals[1:] {
		elem := &Element[T]{
			Data: val,
		}
		curr.Next = elem
		elem.Prev = curr
		curr = elem
	}
	return lst
}
func (l *DoubleLinkedList[T]) AddChildAt(pos int, child *Element[T]) error {
	curr := &l.root
	count := 0
	if pos < 0 {
		return errors.New("position must be non-negative")
	}
	for curr != nil {
		if count == pos-1 {
			curr.Child = child
			return nil
		}
		count++
		curr = curr.Next
	}

	return fmt.Errorf("position must not be out of range. Total elements in the list is %d ", count+1)
}
func (l *DoubleLinkedList[T]) Flatten() error {
	curr := &l.root
	for {
		if curr == nil {
			break
		}
		tmp := curr.Next
		if curr.Child != nil {
			appendChild[T](curr)
			curr = tmp
		} else {
			curr = curr.Next
		}
	}
	return nil
}
func appendChild[T any](curr *Element[T]) {
	tmp := curr.Next
	curr.Next = curr.Child
	curr.Child.Prev = curr
	leaf := flattenChild[T](curr.Child)
	if tmp != nil {
		tmp.Prev = leaf
		leaf.Next = tmp
	}
}
func flattenChild[T any](child *Element[T]) *Element[T] {
	curr := child
	var prev *Element[T]
	for {
		if curr == nil {
			return prev
		}
		if curr.Child != nil {
			appendChild[T](curr)
		}
		prev = curr
		curr = curr.Next
	}
}
func (l *DoubleLinkedList[T]) UnFlatten() error {
	curr := &l.root
	for {
		if curr == nil {
			break
		}
		if curr.Child != nil {
			leaf := getChildLeafNode(curr.Child)
			println(leaf.Data)
			curr.Next = leaf.Next
			if leaf.Next != nil {
				leaf.Next.Prev = curr
			}
		}
		curr = curr.Next
	}
	return nil
}
func getChildLeafNode[T any](child *Element[T]) *Element[T] {
	curr := child
	var nodeWithChild *Element[T]
	prev := curr
	for {
		if curr == nil {
			break
		}
		if curr.Child != nil {
			nodeWithChild = curr.Child
		}
		prev = curr
		curr = curr.Next
	}
	if nodeWithChild == nil {
		return prev
	} else {
		return getChildLeafNode[T](nodeWithChild)
	}
}
func (l *DoubleLinkedList[T]) String() string {
	s := strings.Builder{}
	node := &l.root
	for {
		if node == nil {
			s.WriteString("nil")
			break
		}
		s.WriteString(fmt.Sprintf("%v -> ", node.Data))
		if node.Child != nil {
			s.WriteString(String[T](node.Child))
		}
		node = node.Next
	}
	return s.String()
}
func String[T any](elem *Element[T]) string {
	s := strings.Builder{}
	s.WriteString("[")
	for {
		if elem == nil {
			s.WriteString("]")
			break
		}
		s.WriteString(fmt.Sprintf("%v -> ", elem.Data))
		if elem.Child != nil {
			s.WriteString(String[T](elem.Child))
		}
		elem = elem.Next
	}
	return s.String()
}
func (l *DoubleLinkedList[T]) StringWithoutChild() string {
	s := strings.Builder{}
	node := &l.root
	for {
		if node == nil {
			s.WriteString("nil")
			break
		}
		s.WriteString(fmt.Sprintf("%v -> ", node.Data))
		node = node.Next
	}
	return s.String()
}
