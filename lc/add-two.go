package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val int
	Next *ListNode
}
// A non-recursive solution
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    sumNode := &ListNode{}
	last := sumNode
	var carry int
	for {
		if carry != 0 {
			last.Val += carry
			carry = 0
		}
		if l1 != nil {
			last.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			last.Val += l2.Val
			l2 = l2.Next
		}
		if last.Val > 9 {
			carry = last.Val / 10
			last.Val %= 10
		}

		if l1 == nil && l2 == nil && carry == 0 {
			break
		} else {
			last.Next = &ListNode{}
			last = last.Next
		}
	}
	return sumNode
}

func (listNode *ListNode) String() string {
	if listNode.Next == nil {
		return fmt.Sprint(listNode.Val)
	}
	return listNode.Next.String() + fmt.Sprint(listNode.Val)
}

func AsNodes(digits ...int) *ListNode {
	first := &ListNode{}
	last := first
	for i := 0; i < len(digits); i++ {
		last.Val = digits[i]
		if i + 1  < len(digits) {
			last.Next = &ListNode{}
			last = last.Next
		}
	}

	return first
}

func main(){
	l1 :=AsNodes(2, 4, 3)
	l2 :=AsNodes(5, 6, 4)
	fmt.Printf("%v + %v = %v\n", l1, l2, addTwoNumbers(l1, l2))
}
