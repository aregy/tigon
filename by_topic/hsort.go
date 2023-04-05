package main

import (
	"fmt"
	"os"
	"strconv"
)

func CheckHeap(n []int) {
	for i := len(n) - 1; i > len(n)/2-1; i-- {
		hfied(n, i)
	}
}
func hfied(n []int, head int) {
	m := head
	if len(n) > 2*head+1 && n[m] > n[2*head+1] {
		m = 2*head + 1
	}
	if len(n) > 2*head+2 && n[m] > n[2*head+2] {
		m = 2*head + 2
	}
	if m != head {
		panic(fmt.Errorf("Not heapified because %d > %d in %#v", n[head], n[m], n))
	}
	if head != 0 {
		if n[head] == 94 {
			fmt.Printf("%d: %#v\n", head, n)
		}
		hfied(n, (head-1)/2)
	}
}
func Hsort(n []int) {
	if len(n) == 0 {
		return
	}
	for i := len(n) - 1; i > len(n)/2-1; i-- {
		rtBubble(n[:i+1])
	}
	Hsort(n[1:])
}
func rtBubble(n []int) {
	rt := len(n) - 1
	for rt != 0 {
		p := (rt - 1) / 2
		if rt == p {
			break
		}
		if n[p] > n[rt] {
			n[p], n[rt] = n[rt], n[p]
		}
		rt = p
	}
}

func getNums() []int {
	s := os.Args[1:]
	n := make([]int, len(s))
	for i, v := range s {
		var err error
		n[i], err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
	}
	return n
}

func main() {
	n := getNums()
	fmt.Println(n)
	//Hsort(n)
	CheckHeap(n)
	fmt.Println(n)
}
