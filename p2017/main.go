package main

import "fmt"

func gridGame(grid [][]int) int64 {
	if len(grid[0]) < 2 {
		return 0
	}
	for idx := len(grid[0]) - 2; idx > -1; idx-- {
		grid[0][idx] += grid[0][idx+1]
	}
	for idx := range grid[1] {
		if idx == 0 {
			continue
		}
		grid[1][idx] += grid[1][idx-1]
	}
	values := make([]int, 0, len(grid[0]))
	values = append(values, grid[0][1])
	for idx := 1; idx <= len(grid[0])-2; idx++ {
		values = append(values, max(grid[0][idx+1], grid[1][idx-1]))
	}
	values = append(values, grid[1][len(grid[1])-2])
	min := values[0]

	for idx := 1; idx < len(values); idx++ {
		if values[idx] < min {
			min = values[idx]
		}
	}
	return int64(min)
}

func main() {
	testcases := [][][]int{{{2, 5, 4}, {1, 5, 1}}}
	testcases = append(testcases, [][]int{{3, 3, 1}, {8, 5, 2}})
	for idx, test := range testcases {
		fmt.Printf("Testcase #%d: given %v...", idx+1, test)
		fmt.Print(gridGame(test))
	}
}
