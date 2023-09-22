package main

import (
	"fmt"
	"os"
)

func main() {
	words := os.Args[1:]
	for _, w := range words {
		res := []rune(w)
		for i := 0; i < len(res) / 2; i++ {
			res[len(res)-i-1], res[i] = res[i], res[len(res)-1-i]
		}
		fmt.Println(string(res))
	}
}
