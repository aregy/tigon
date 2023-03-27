package main

import(
	"os"
	"strconv"
	"fmt"
)

func main(){
	nums := make([]int, len(os.Args[1:]))
	for i, arg := range os.Args[1:] {
		var x, err = strconv.Atoi(arg)
		if err != nil {
			panic(err)
		}
		nums[i] = x
	}
	fmt.Println(nums, max(nums...), min(nums...), minFree(nums...))
}

func min(n ...int) int{
	if len(n) == 0 {
		panic("Empty set has no min element")
	}
	min := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] < min {
			min = n[i]
		}
	}
	return min
}
func max(n ...int) int{
	if len(n) == 0 {
		panic("Empty set has no max element")
	}
	max := n[0]
	for i := 1; i < len(n); i++ {
		if n[i] > max {
			max = n[i]
		}
	}
	return max
}
func minFree(n ...int) int{
	size := len(n)
	free := make([]*struct{}, size+1)
	for _, a := range n {
		if a <= size {
			free[a] = &struct{}{}
		}
	}
	for i := 0; i < size+1; i++ {
		if free[i] == nil {
			return i
		}
	}
	return size
}
