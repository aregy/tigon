package main

import (
	"fmt"
	"os"
	"strconv"
)

func insertion_sort(args []int) {
	ops := 0
	for i := 1; i < len(args); i++ {
		if args[i] < args[i-1] {
			args[i], args[i-1] =
				args[i-1], args[i]

		}
		ops++
	}
	fmt.Printf("%#v (%d)\n\tsorted in  %d comparisons\n",
		args,
		len(args),
		ops)
}

func main() {
	args := make([]int, 0, len(os.Args[1:]))
	for _, el := range os.Args[1:] {
		if k, err := strconv.Atoi(string(el)); err == nil {
			args = append(args, k)
		}
	}
	insertion_sort(args)
}
