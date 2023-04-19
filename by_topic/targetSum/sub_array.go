package main

import "fmt"

func force_findSubarray(src []int, targetSum int) []int {
	var tot int
	for left := 0; left < len(src); left++ {
		tot = 0
		for right := left; right < len(src); right++ {
			tot += src[right]
			if tot == targetSum {
				return src[left : right+1]
			}
		}
	}
	return nil
}

func findSubarray(src []int, targetSum int) []int {
	areas := make(map[int]int)
	sum := 0
	areas[0] = -1
	for i, v := range src {
		sum += v
		areas[sum] = i
		if complement, exists := areas[sum-targetSum]; exists {
			return src[complement+1 : i+1]
		}
	}
	return nil
}

func main() {
	s := []int{1, 2, 4, 5, 8, 9, 10}
	var h func() = func() {
		// TODO handle ^Csignal: interrupt
		recover()
		fmt.Println("Good bye...")
	}
	defer h()
	var t int
	for {
		fmt.Printf("%v(+): ", s)
		fmt.Scanf("%d", &t)
		fmt.Printf(" O(n^2) :: %v\n", force_findSubarray(s, t))
		fmt.Printf("  O(n)  :: %v\n\n", findSubarray(s, t))
	}
}
