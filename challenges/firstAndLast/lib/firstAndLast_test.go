package main

import (
	"math/rand"
	"testing"
)

// func TestFirstAndLast(t *testing.T) {
func main() {
	inputSize := 30
	ceiling := 100
	input := make([]int, inputSize)
	fmt.Println(input)
	for len(input) < inputSize {
		next := rand.Intn(ceiling)
		if next >= input[len(input)-1] {
			input[len(input)] = next
		}
	}
}
