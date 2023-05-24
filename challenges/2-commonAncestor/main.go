package main

import "fmt"

type Node struct {
	Val    rune
	Left   *Node
	Right  *Node
	Parent *Node
}

func (n *Node) FirstCommonAncestor(descendents ...*Node) *Node {
	panic("Unimplemented")
}

func main() {
	root := Node{}
	root.Val = 'a'
	root.Left = &Node{}
	root.Left.Val = 'b'
	root.Left.Parent = &root
	root.Right = &Node{}
	root.Right.Val = 'c'
	root.Right.Parent = &root
	root.Right.Left = &Node{}
	root.Right.Val = 'd'
	a := root.Right.Left
	root.Right.Left.Parent = root.Right
	root.Right.Right = &Node{}
	root.Right.Right.Val = 'e'
	b := root.Right.Right
	root.Right.Right.Parent = root.Right
	fmt.Println(root.FirstCommonAncestor(a, b))
}
