package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	MaxInt int = int(^uint(0) >> 1)
	MinInt int = -int(^uint(0)>>1) - 1
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func (node *Node) String() string {
	if node == nil {
		return "-"
	}
	return node.StringHelper(getDepth(node, 0))
}

func (node *Node) StringHelper(n int) string {
	if node == nil {
		return " "
	}
	var repr string
	repr += strings.Repeat("  ", min(int(math.Exp2(float64(n))), 25*n/10))
	repr += fmt.Sprint(node.Val)
	repr += "\n"
	repr += leftRightConcat(node.Left.StringHelper(n-1), node.Right.StringHelper(n-1))
	return repr
}

func IsBalancedTree(node *Node, interval ...int) bool {
	if node == nil {
		return true
	}
	if len(interval) >= 1 && interval[0] > node.Val {
		fmt.Println(interval[0], "≮", node.Val)
		return false
	}
	if len(interval) >= 2 && interval[1] < node.Val {
		fmt.Println(interval[1], "≯", node.Val)
		return false
	}
	lmin, rmax := MinInt, MaxInt
	if len(interval) >= 1 {
		lmin = interval[0]
	}
	if len(interval) >= 2 {
		rmax = interval[1]
	}

	return IsBalancedTree(node.Left, lmin, min(node.Val, rmax)) && IsBalancedTree(node.Right, max(node.Val, lmin), rmax)
}
func min(n ...int) int {
	if len(n) == 0 {
		panic("Max entry not present in empty set")
	}
	min := n[0]
	for i, _ := range n {
		if n[i] < min {
			min = n[i]
		}
	}
	return min
}
func max(n ...int) int {
	if len(n) == 0 {
		panic("Max entry not present in empty set")
	}
	max := n[0]
	for i, _ := range n {
		if n[i] > max {
			max = n[i]
		}
	}
	return max
}

func leftRightConcat(left, right string) string {
	leftEntries := strings.Split(left, "\n")
	rightEntries := strings.Split(right, "\n")
	result := make([]string, max(len(leftEntries), len(rightEntries)))
	for i := 0; i < len(result); i++ {
		if i < len(leftEntries) {
			result[i] += leftEntries[i]
		}
		if i < len(leftEntries) && i < len(rightEntries) {
			result[i] += "  "
		}
		if i < len(rightEntries) {
			result[i] += rightEntries[i]
		}
	}
	return strings.Join(result, "\n")
}

func getDepth(node *Node, level int) int {
	left, right := level, level
	if node.Left != nil {
		left = getDepth(node.Left, level+1)
	}
	if node.Right != nil {
		right = getDepth(node.Right, level+1)
	}

	if left > right {
		return left
	}
	return right
}

func generateMessage(t *Node) {
	if IsBalancedTree(t) {
		fmt.Printf("%v is balanced\n", t)
	} else {
		fmt.Printf("%v is unbalanced\n.", t)
	}
}

func main() {
	var t Node = Node{1, &Node{11, &Node{22, &Node{333, nil, nil}, &Node{444, nil, nil}}, nil}, &Node{33, &Node{444, nil, nil}, &Node{444, nil, nil}}}
	generateMessage(&t)
	t = Node{8, &Node{6, &Node{4, &Node{3, nil, nil}, &Node{5, nil, nil}}, nil}, &Node{11, &Node{10, nil, nil}, &Node{12, nil, nil}}}
	generateMessage(&t)
}
