package main

import "fmt"

type Node struct {
	Val  any
	Next *Node
}

func (node *Node) String() string {
	if node.Next != nil {
		return fmt.Sprintf("%d -> %v", node.Val, node.Next)
	}
	return fmt.Sprintln(node.Val)
}

func GetNums(n int, k int) *Node {
	index := 0
	initFn := func(node *Node) {
		node.Val = index + 1
		index++
	}
	first := &Node{}
	current := first
	for i := 0; ; i++ {
		initFn(current)
		if i == n-1 {
			break
		}
		current.Next = &Node{}
		current = current.Next
	}
	return first
}
func Reversed(n *Node) *Node {

	var prev *Node = nil
	next := n.Next
	if next == nil {
		return n
	}
	for {
		next = n.Next
		n.Next = prev
		if next == nil {
			break
		}
		prev = n
		n = next
	}
	return n
}
func main() {
	nums := GetNums(10, 1)
	fmt.Println(nums)
	nums = Reversed(nums)
	fmt.Println(nums)
}
