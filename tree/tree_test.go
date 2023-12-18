package main

import "testing"

func TestTree(t *testing.T) {
	tree := NewTree(3)
	n := tree.AddNode(2)
	n.PrintNode()
	tree.AddNode(1)
	tree.AddNode(4)
	tree.AddNode(5)
	tree.AddNode(7)
	tree.AddNode(8)
	tree.AddNode(9)
	tree.AddNode(4)
	tree.AddNode(3)
	n = tree.FindNode(10)
	if n != nil {
		t.FailNow()
	}
	n = tree.FindNode(8)
	if n == nil {
		t.FailNow()
	}
}
func TestTraverse(t *testing.T) {
	tree := NewTree(3)
	tree.AddNode(2)
	tree.AddNode(1)
	tree.AddNode(4)
	tree.AddNode(5)
	tree.AddNode(7)
	tree.AddNode(8)
	tree.AddNode(9)
	tree.AddNode(4)
	tree.AddNode(3)
	TraversePreOrder(&tree.root)
	println()
	TraversePreOrderLoop(&tree.root)
	println()
	TraverseInOrder(&tree.root)
	println()
	TraversePostOrder(&tree.root)
}
func TestCommonAncestor(t *testing.T) {
	tree := NewTree(3)
	tree.AddNode(2)
	tree.AddNode(1)
	tree.AddNode(4)
	tree.AddNode(5)
	tree.AddNode(7)
	tree.AddNode(8)
	tree.AddNode(9)
	tree.AddNode(4)
	tree.AddNode(3)
	n := FindCommonAncestor(&tree.root, 3, 9)
	println(n.Data)
}

func TestTreeHeight(t *testing.T) {
	tree := NewTree(3)
	tree.AddNode(2)
	tree.AddNode(1)
	tree.AddNode(4)
	tree.AddNode(5)
	tree.AddNode(7)
	tree.AddNode(8)
	tree.AddNode(9)
	tree.AddNode(4)
	tree.AddNode(3)
	height := TreeHeight(&tree.root)
	t.Log(height)
}
func TestHeapifyTree(t *testing.T) {
	tree := NewTree(3)
	tree.AddNode(2)
	tree.AddNode(1)
	tree.AddNode(4)
	tree.AddNode(5)
	tree.AddNode(6)
	tree.AddNode(7)
	tree.AddNode(8)
	tree.AddNode(9)
	n := HeapifyBinaryTree(&tree.root)
	TraversePreOrder(n)
}
