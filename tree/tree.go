package main

import (
	"fmt"
	"math"
	"sort"
)

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}
type Tree struct {
	root Node
}

func (n *Node) PrintNode() {
	fmt.Printf("\t%d\t\n", n.Data)
	if n.Left != nil {
		fmt.Printf("%d", n.Left.Data)
	}
	if n.Right != nil {
		fmt.Printf("\t\t%d", n.Right.Data)
	}
	println()
}
func NewTree(data int) *Tree {
	rootNode := Node{
		Data: data,
	}
	return &Tree{
		root: rootNode,
	}
}
func (n *Node) AddChild(data int) *Node {
	node := &Node{
		Data: data,
	}
	if data < n.Data {
		if n.Left == nil {
			n.Left = node
		} else {
			n.Left.AddChild(data)
		}
	} else {
		if n.Right == nil {
			n.Right = node
		} else {
			n.Right.AddChild(data)
		}
	}
	return node
}
func (t *Tree) FindNode(data int) *Node {
	var findNode func(n *Node) *Node
	findNode = func(n *Node) *Node {
		if n == nil {
			return nil
		}
		if n.Data == data {
			println("data found")
			return n
		} else if data < n.Data {
			return findNode(n.Left)
		} else {
			return findNode(n.Right)
		}
	}
	return findNode(&t.root)
}
func (t *Tree) AddNode(data int) *Node {
	return t.root.AddChild(data)
}
func TraversePreOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Printf("%d ", n.Data)
	if n.Left != nil {
		TraversePreOrder(n.Left)
	}
	if n.Right != nil {
		TraversePreOrder(n.Right)
	}
}
func TraversePreOrderLoop(n *Node) {
	if n == nil {
		return
	}
	stack := []*Node{n}
	for len(stack) != 0 {
		len := len(stack)
		n := stack[len-1]
		stack = stack[:len-1]
		fmt.Printf("%d ", n.Data)
		if n.Right != nil {
			stack = append(stack, n.Right)
		}
		if n.Left != nil {
			stack = append(stack, n.Left)
		}

	}
}
func TraverseInOrder(n *Node) {
	if n == nil {
		return
	}
	if n.Left != nil {
		TraversePreOrder(n.Left)
	}
	fmt.Printf("%d ", n.Data)
	if n.Right != nil {
		TraversePreOrder(n.Right)
	}
}
func TraversePostOrder(n *Node) {
	if n == nil {
		return
	}
	if n.Left != nil {
		TraversePreOrder(n.Left)
	}
	if n.Right != nil {
		TraversePreOrder(n.Right)
	}
	fmt.Printf("%d ", n.Data)
}

func TreeHeight(n *Node) int {
	if n == nil {
		return 0
	}
	return 1 + int(math.Max(float64(TreeHeight(n.Left)), float64(TreeHeight(n.Right))))
}

func FindCommonAncestor(n *Node, val1, val2 int) *Node {
	for n != nil {
		if val1 < n.Data && val2 < n.Data {
			n = n.Left
		} else if val1 >= n.Data && val2 >= n.Data {
			n = n.Right
		} else {
			return n
		}
	}
	return nil
}
func traverse(n *Node, arr *[]*Node) {
	if n == nil {
		return
	}
	*arr = append((*arr), n)
	traverse(n.Left, arr)
	traverse(n.Right, arr)
}
func HeapifyBinaryTree(n *Node) *Node {
	arr := []*Node{}
	traverse(n, &arr)
	size := len(arr)
	sort.SliceStable(arr, func(i, j int) bool { return arr[i].Data < arr[j].Data })
	for i := 0; i < size; i++ {
		n := arr[i]
		left := 2*i + 1
		right := left + 1
		if left < size {
			n.Left = arr[left]
		}
		if right < size {
			n.Right = arr[right]
		}
		n.PrintNode()
	}
	return arr[0]
}
