package uniq

import "fmt"

func FindUniq(nums []int) int {
	vals := make(map[int]struct{}, len(nums)/2+1)
	for _, v := range nums {
		if _, exists := vals[v]; exists {
			delete(vals, v)
			continue
		}
		vals[v] = struct{}{}
	}
	fmt.Println(nums, vals)
	for key, _ := range vals {
		return key
	}
	panic(fmt.Sprintf("No uniq values exist in %v", vals))
}
