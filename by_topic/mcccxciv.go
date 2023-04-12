package main

import "fmt"

type Node struct {
	Val  int
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
func Sort(n *Node) *Node {
	arr := []*Node{}
	for n != nil {
		arr = append(arr, n)
		n = n.Next
	}
	partition(arr, 0, len(arr)-1)
	for i, k := range arr {
		if i != len(arr)-1 {
			k.Next = arr[i+1]
		} else {
			k.Next = nil
		}
	}
	return arr[0]
}
func partition(a []*Node, l, r int) {
	if r-l < 2 {
		return
	}
	index := l - 1
	pivot := a[r].Val
	for i := l; i < r; i++ {
		if a[i].Val < pivot {
			index++
			a[i], a[index] = a[index], a[i]
		}
	}
	index++
	a[r], a[index] = a[index], a[r]
	partition(a, l, index-1)
	partition(a, index+1, r)
}
func main() {
	nums := GetNums(10, 1)
	fmt.Println(nums)
	nums = Reversed(nums)
	fmt.Println(nums)
	nums = Sort(nums)
	fmt.Println(nums)
}
