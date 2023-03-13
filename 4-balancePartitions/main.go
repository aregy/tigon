package main

import (
	"fmt"
	"os"
	"strconv"
)

type subsequence []int

func (s subsequence) Sum() int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}

func (s subsequence) Min() int {
	var min int
	for i, v := range s {
		if i == 0 {
			min = i
		}
		if v < s[min] {
			min = i
		}
	}
	return min
}

func ToLists(args []string) []subsequence {
	parts := []subsequence{subsequence{}, subsequence{}, subsequence{}}
	var i int
	for _, v := range args {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		parts[i] = append(parts[i], num)
		i++
		if i == len(parts) {
			i = 0
		}
	}
	return parts
}

func main() {
	partitions := ToLists(os.Args[1:])
	fmt.Println(partitions)
}
