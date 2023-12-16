package main

import "testing"

func TestAddChildAtPos(t *testing.T) {
	l := CreateDoubleLinkedList[int]([]int{1, 2, 3, 4, 5})
	t.Log(l.String())
	child5 := CreateDoubleLinkedList[int]([]int{51, 52, 53, 54, 55})
	l.AddChildAt(5, &child5.root)
	child3 := CreateDoubleLinkedList[int]([]int{31, 32, 33, 34, 35})
	l.AddChildAt(3, &child3.root)
	child3Level2 := CreateDoubleLinkedList[int]([]int{311, 312, 313, 314, 315})
	child3.AddChildAt(1, &child3Level2.root)
	child311Level3 := CreateDoubleLinkedList[int]([]int{351, 352, 353, 354, 355})
	child3.AddChildAt(5, &child311Level3.root)
	t.Log(l.String())
	l.Flatten()
	t.Log(l.StringWithoutChild())
	l.UnFlatten()
	t.Log(l.StringWithoutChild())
}
