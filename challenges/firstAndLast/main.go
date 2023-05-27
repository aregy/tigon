package main

import (
	"fmt"
	"math/rand"
)

// func TestFirstAndLast(t *testing.T) {
func main() {
	inputSize := 30
	height := 6
	input := make([]int, 0, inputSize)
	floor := 0
	org := 0
	first := -1
	last := 0
	count := 0
	for org != 0 && len(input) < inputSize {
		next := rand.Intn(height) + floor
		floor = next
		if next == input[len(input)-1] {
			count++
			org = next
			if first == -1 {
				first = len(input) - 1
			}
		} else {
			last = first + count
		}
		input = append(input, next)
	}
	fmt.Println(input, org, first, last)
}
