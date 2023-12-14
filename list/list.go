package main

import (
	"fmt"
	"strings"
)

type ListElement[T any] struct {
	Next *ListElement[T]
	Data T
}
type List[T any] struct {
	root ListElement[T]
	len  int
}

func (l *List[T]) Len() int {
	return l.len
}

func (l *List[T]) PushFront(value T) *ListElement[T] {
	elem := ListElement[T]{
		Data: value,
	}
	if l.len != 0 {
		//change root to new element
		temp := l.root
		elem.Next = &temp
	}
	l.root = elem
	l.len++
	return &elem
}
func (l *List[T]) PushBack(value T) *ListElement[T] {
	elem := ListElement[T]{
		Data: value,
	}
	if l.len == 0 {
		l.root = elem
	} else {
		node := &l.root
		for {
			if node.Next == nil {
				node.Next = &elem
				break
			}
			node = node.Next
		}
	}
	l.len++
	return &elem
}
func (l *List[T]) FindFirst(predicate func(elem T) bool) *ListElement[T] {
	node := &l.root
	for {
		if predicate(node.Data) {
			return node
		}
		if node.Next == nil {
			break
		}
		node = node.Next
	}
	return nil
}
func NewListFromSlice[T any](slice []T) *List[T] {
	lst := &List[T]{}
	if len(slice) == 0 {
		return lst
	}
	lst.root = ListElement[T]{
		Data: slice[0],
	}
	lst.len = 1
	curr := &lst.root
	for _, val := range slice[1:] {
		elem := &ListElement[T]{
			Data: val,
		}
		curr.Next = elem
		curr = elem
		lst.len++
	}
	return lst
}
func (l *List[T]) DeleteElement(elem *ListElement[T]) bool {
	//case 1 -> removing head
	//case 2 -> removing the last element
	//case 3 -> list is empty
	//case 4 -> removing in the intermediate
	if l.len == 0 {
		return false
	}
	curr := &l.root
	var prev *ListElement[T] = nil
	for {
		if curr == nil {
			return false
		}
		if curr == elem {
			if prev == nil {
				//removing head
				l.root = *elem.Next
			} else if curr.Next == nil {
				//removing tail
				prev.Next = nil
			} else {
				prev.Next = elem.Next
			}
			elem = nil
			l.len--
			return true
		}
		prev = curr
		curr = curr.Next
	}
}
func (l *List[T]) String() string {
	s := strings.Builder{}
	node := &l.root
	for {
		s.WriteString(fmt.Sprintf("%v -> ", node.Data))
		if node.Next == nil {
			s.WriteString("nil")
			break
		}
		node = node.Next
	}
	return s.String()
}

func main() {
	l := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	l.PushFront(0)
	l.PushBack(9)
	println(l.len, l.String())
	elem := l.FindFirst(func(elem int) bool { return elem == 5 })
	println(elem.Data, elem.Next.Data)
}
