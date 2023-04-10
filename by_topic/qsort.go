package main

import "fmt"

func partition(nums []int, l int, r int) {
	if r-l < 1 {
		return
	}
	pivot := nums[r]
	index := l - 1
	for i := l; i < r; i++ {
		if nums[i] < pivot {
			nums[index+1], nums[i] = nums[i], nums[index+1]
			index++
		}
	}
	nums[index+1], nums[r] = nums[r], nums[index+1]
	index++

	partition(nums, l, index-1)
	partition(nums, index+1, r)
}

func qsort(nums []int) {
	partition(nums, 0, len(nums)-1)
}

func stat(nums []int) {
	fmt.Println(nums)
	qsort(nums)
	fmt.Println(nums)
}

func main() {
	stat([]int{4, 4, 9, 1, 2, -3, 3, 10, 11, 8})
}
