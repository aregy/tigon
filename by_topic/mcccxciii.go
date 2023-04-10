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

func JosephusSeq(n int, k int) {
	index := 0
	initFn := func(node *Node) {
		node.Val = index + 1
		index++
	}
	nums := GetLinkedList(n, initFn)
	fmt.Println(nums)
	// TODO
}
func GetLinkedList(n int, initFn func(*Node)) *Node {
	first := &Node{}
	current := first
	for i := 0; ; i++ {
		current.Val = i + 1
		if i == n-1 {
			break
		}
		current.Next = &Node{}
		current = current.Next
	}
	return first
}

func main() {
	JosephusSeq(10, 1)
}
