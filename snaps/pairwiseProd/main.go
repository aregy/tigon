package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
const green string = "\033[1;34m%s\033[0m"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	nums := make([]int, 0, 1)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			break
		}
		nums = append(nums, n)
	}
	fmt.Println("Given", nums)
	fmt.Printf(green, strconv.Itoa(maxProduct(nums)))
	fmt.Println()
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
	p1, p2 := max*max2, min*min2
	if p1 > p2 {
		return p1
	}
	return p2
}
