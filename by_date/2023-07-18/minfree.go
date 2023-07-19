package main

import "fmt"

func main() {
	n := []int{3, 9, 0, 1}
	fmt.Println(n, minfree(n))
}

func minfree(nums []int) int {
	values := make(map[int]struct{})
	for _, el := range nums {
		values[el] = struct{}{}
	}
	k := 0
	for ; ; k++ {
		if _, ok := values[k]; ok == false {
			break
		}
	}
	return k
}
