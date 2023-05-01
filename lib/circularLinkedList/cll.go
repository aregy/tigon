package circ

import (
	"fmt"
	"strings"
)

type CircularLinkedList struct {
	Length int
	Tail   *Node
}

type Node struct {
	Val  any
	Next *Node
}

func (cll *CircularLinkedList) String() string {
	if cll.Length == 0 {
		return "<EMPTY>"
	}
	node := cll.Tail
	repr := []string{}
	for i := 0; i < cll.Length; i++ {
		node = node.Next
		switch v := node.Val.(type) {
		case string:
			repr = append(repr, v)
		case int:
			repr = append(repr, fmt.Sprint(v))
		}

	}
	return strings.Join(repr, "->")
}
