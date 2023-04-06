package main

import (
	"fmt"
)

func AbsSort(nums []int) {
	var hfyDown func(int)
	hfyDown = func(j int) {
		m := j
		l := 2*m + 1
		r := 2*m + 2
		if l < len(nums) && abs(nums[l]) > abs(nums[m]) {
			m = l
		}
		if r < len(nums) && abs(nums[r]) > abs(nums[m]) {
			m = r
		}
		if m != j {
			nums[m], nums[j] = nums[j], nums[m]
			hfyDown(m)
		}
	}
	for i := len(nums)/2 - 1; i >= 0; i-- {
		hfyDown(i)
	}
	for i := len(nums) - 1; i > 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		hfyDown(0)
	}
}
func TriplePoint(nums []int) int {
	//for i := 0; i < len(nums)-1; i++ {
	AbsSort(nums)
	//}
	i, j, k := 0, 1, 2
	if arity(nums[i], nums[j], nums[k])%2 == 0 {
		return product(nums[i], nums[j], nums[k])
	}
	pos := make([]int, 0, len(nums))
	npos := make([]int, 0, len(nums))
	for _, v := range nums {
		if v > 0 {
			pos = append(pos, v)
		} else {
			npos = append(npos, v)
		}
	}
	if len(pos) >= 3 {
		return product(pos[0], pos[1], pos[2])
	} else if len(pos) == 2 {
		return product(pos[0], pos[1], max(npos...))
	} else if len(pos) == 1 {
		return product(pos[0], npos[0], npos[1])
	}
	return product(npos[0], npos[1], npos[2])
}

func product(n ...int) int {
	if len(n) == 0 {
		panic(fmt.Errorf("Empty arguments in prodcut(n ...int) call\n"))
	}
	prod := 1
	for _, v := range n {
		prod *= v
	}
	return prod
}
func arity(n ...int) int {
	if len(n) == 0 {
		return 0
	}
	ar := 0
	for _, v := range n {
		if v < 0 {
			ar += 1
		}
	}
	return ar
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(n ...int) int {
	if len(n) == 0 {
		panic(fmt.Errorf("Empty set contains not minimum element"))
	}
	m := n[0]
	for _, v := range n {
		if v < m {
			m = v
		}
	}
	return m
}
func max(n ...int) int {
	if len(n) == 0 {
		panic(fmt.Errorf("Empty set contains not minimum element"))
	}
	m := n[0]
	for _, v := range n {
		if v > m {
			m = v
		}
	}
	return m
}
func stat(inputs []int) {
	fmt.Println(inputs)
	fmt.Println(TriplePoint(inputs))
}
func main() {
	stat([]int{1, 4, 10, 9, -9, 5, 0}) // 450
	stat([]int{-4, -4, 2, 8})          // 128
}
