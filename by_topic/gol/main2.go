package main

import (
	"fmt"
	"os"
	"strings"
)

func reverseWordOrder(words ...string) string {
	for i := 0; i < len(words)/2; i++ {
		words[i], words[len(words)-1-i] = words[len(words)-1-i], words[i]
	}
	return strings.Join(words, " ")
}

func main() {
	args := os.Args[1:]
	fmt.Println(reverseWordOrder(args...))
	var reversed []rune
	for _, w := range args {
		res := []rune(w)
		for i := 0; i < len(res)/2; i++ {
			res[len(res)-i-1], res[i] = res[i], res[len(res)-1-i]
		}
		reversed = append(reversed, res...)
		reversed = append(reversed, rune(' '))
	}
	fmt.Printf("%v\n", string(reversed))
}
