package main

import (
	"os"
	"fmt"
	"strconv"
)

func TwoSum(target int, nums []int) (int, int) {
	complement := make([]int, len(nums))
	copy(complement, nums)
	fmt.Println(complement)
	return 0, len(complement) -1
}
func parseList(args []string) []int{
	nums := make([]int, len(args))
	for i, v := range args {
		var err error
		nums[i], err = strconv.Atoi(v)
		if err != nil {
			panic(fmt.Errorf("%s in %#v does not have an integer conversion\n", v, args))
		}
	}
	return nums
}
func isValidInput(nums []int) bool {
	vals := make(map[int]any)
	for _, val := range nums[1:] {
		vals[val] = struct{}{}
	}
	for i := 0; i < nums[0]; i++ {
		if vals[i] != nil && vals[nums[0] - i] != nil {
			return true
		}
	}
	return false
}
func main(){
	nums := parseList(os.Args[1:])
	if !isValidInput(nums) {
		panic(fmt.Errorf("%d is not a valid target", nums[0]))
	}
	TwoSum(nums[0], nums[1:])
}
