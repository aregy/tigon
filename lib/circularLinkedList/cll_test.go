package circ

import (
	"testing"
)

func TestCircularLinkedList(t *testing.T) {
	//c := CircularLinkedList{}

}

func TestCircularLinkedList_String(t *testing.T) {
	//got := t
	n1 := &Node{}
	n2 := &Node{}
	n3 := &Node{}
	n1.Val = 1
	n2.Val = 2
	n3.Val = 3
	n1.Next = n2
	n2.Next = n3
	n3.Next = n1
	c := CircularLinkedList{}
	c.Tail = n3
	c.Length = 3
	want := "1->2->3"
	got := c.String()
	if got != want {
		t.Errorf("Got '%s' but expected '%s'", got, want)
	}
}
