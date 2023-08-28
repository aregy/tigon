package main

import (
	"fmt"
)

func main() {
	list := []int{900000, 10000}
	fmt.Println(maxProduct(list))
}

func maxProduct(nums []int) int {
	var max, max2 int
	var min, min2 int
	for i, el := range nums {
		switch true {
		case i == 0:
			max = el
		case el > max:
			max2 = max
			max = el
		case el > max2:
			max2 = el
		}
		switch true {
		case i == 0:
			min = el
		case el < min:
			min2 = min
			min = el
		case el < min2:
			min2 = el
		}
	}
	p1, p2 := max * max2, min * min2
	if p1 > p2{
		return p1
	}
	return p2
}
