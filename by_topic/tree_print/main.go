package main

import "strings"

type Node struct {
	Val   any
	Left  *Node
	Right *Node
}

func (n *Node) String() string {
	if n.Val == nil {
		return "-"
	}
	return string(n.Val)
}

func PrintTree(n *Node) {
	levels := map[int]string{}
	PrintLevel(n, &levels, l)
	return strings.Join(levels)
}

func PrintLevel(n *Node, levels map[int]string, l level) {
	if s, ok := levels[l]; ok {
		(*levels)[l] = s + " " + string(n.Val)
	} else {
		(*levels)[l] = string(n.Val)
	}
	if n.Left != nil {
		PrintLevel(n, levels, l+1)
	}
	if n.Right != nil {
		PrintLevel(n, levels, l+1)
	}
}

func main() {
	PrintTree(*Node{1,
		*Node{2,
			*Node{4, nil, nil},
		},
		*Node{3, nil, nil},
	})
}
