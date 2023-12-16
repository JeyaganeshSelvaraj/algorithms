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

func TestMthToLastElement(t *testing.T) {
	l := NewListFromSlice[int]([]int{})
	if _, err := l.FindMthToLastElement(1); err == nil {
		t.FailNow()
	}
	l.PushFront(1)
	if elem, _ := l.FindMthToLastElement(0); elem == nil || elem.Data != 1 {
		t.FailNow()
	}
	l = NewListFromSlice[int]([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	if elem, _ := l.FindMthToLastElement(0); elem == nil || elem.Data != 9 {
		t.FailNow()
	}
	if elem, _ := l.FindMthToLastElement(1); elem == nil || elem.Data != 8 {
		t.FailNow()
	}
	if elem, _ := l.FindMthToLastElement(5); elem == nil || elem.Data != 4 {
		t.FailNow()
	}
}
func TestCyclicList(t *testing.T) {
	l := NewListFromSlice([]int{1, 2, 3, 4, 5, 6, 7})
	if l.IsCyclic() {
		t.Error("List is not cyclic")
	}
	elem := l.FindFirst(func(elem int) bool { return elem == 5 })
	l.InsertAfter(6, elem)
	if !l.IsCyclic() {
		t.Error("List is cyclic")
	}
}
