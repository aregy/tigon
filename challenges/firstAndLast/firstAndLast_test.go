package firstAndLast

import "math/rand"

func TestFirstAndLast(t *testing.T) {
	inputSize = 30
	ceiling = 200
	input := make([]int, inputSize)
	while len(i) < 30 {
		next := rand.Intn(ceiling)
		if next >= input[len(input)-1] {
			input[len(input)] = next
		}
	}

}
