package main

import (
	"os"
	"fmt"
	"strconv"
)

func TwoSum(target int, nums []int) (int, int) {
	complement := make([]int, len(nums))
	copy(complement, nums)
	for i := 0; i < len(complement); i++ {
		complement[i] = -nums[i] + target
	}
	vals := make(map[int]int, len(nums))

	for i, v := range nums {
		cv := target - v
		if _, ok := vals[cv]; !ok {
			vals[v] = i
		} else if i > vals[cv] {
				return vals[cv], i
			}	else if vals[cv] >= i {
			return i, vals[cv]
		}
	}
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
	fmt.Println(TwoSum(nums[0], nums[1:]))
}
