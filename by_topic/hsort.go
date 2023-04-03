package main

import "fmt"

func Hsort2(n []int) {
	if len(n) == 0 {
		return
	}
	Hsort(n[1:])

	for i := 0; i < len(n); i++ {
		heap(n[:i+1])
	}

}
func Hsort(n []int) {
	if len(n) == 0 {
		return
	}
	for i := 0; i < len(n); i++ {
		heap(n[:i+1])
	}
	Hsort(n[1:])
}
func heap(n []int) {
	rt := len(n) - 1
	for {
		p := (rt - 1) / 2
		if rt == p {
			break
		}
		if n[p] > n[rt] {
			fmt.Printf("%d <- %d\n", n[p], n[rt])
			n[p], n[rt] = n[rt], n[p]
		} else {
			break
		}
		rt = p
	}
}

func main() {
	n := []int{5, 1, 2, 9, 3, 9, 4}
	fmt.Println(n)
	Hsort(n)
	fmt.Println(n)
}
