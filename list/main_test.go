package main

import "testing"

func TestFindFirst(t *testing.T) {
	l := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	elem := l.FindFirst(func(elem int) bool { return elem == 5 })
	if elem == nil {
		t.FailNow()
	}
}
func TestDeleteIntermediate(t *testing.T) {
	l := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	elem := l.FindFirst(func(elem int) bool { return elem == 5 })
	res := l.DeleteElement(elem)
	if !res {
		t.FailNow()
	}
	elem = l.FindFirst(func(elem int) bool { return elem == 5 })
	if elem != nil {
		t.FailNow()
	}
}
func TestDeleteHead(t *testing.T) {
	l := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	elem := l.FindFirst(func(elem int) bool { return elem == 1 })
	res := l.DeleteElement(elem)
	if !res {
		t.FailNow()
	}
	elem = l.FindFirst(func(elem int) bool { return elem == 1 })
	if elem != nil {
		t.FailNow()
	}
}
func TestDeleteTail(t *testing.T) {
	l := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	elem := l.FindFirst(func(elem int) bool { return elem == 7 })
	res := l.DeleteElement(elem)
	if !res {
		t.FailNow()
	}
	t.Log(l.String())
	elem = l.FindFirst(func(elem int) bool { return elem == 7 })
	if elem != nil {
		t.FailNow()
	}
}
