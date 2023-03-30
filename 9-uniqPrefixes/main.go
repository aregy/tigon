package main

import (
	"fmt"
)

func uniqPrefixes(words []string) {

	for i := 0; i < len(words); i++ {
		width := 1

		for j := 0; j < len(words); j++ {
			if i == j {
				continue
			}
			for k := 0; k < len(words[i]); k++ {
				if words[i][:width] == words[j][:width]{
					width += 1
				}
			}
		}
		words[i] = words[i][:width]
	}
}

func main() {

	v1 := []string{"joma", "john", "jack", "techlead"}
	fmt.Println(v1)
	uniqPrefixes(v1)
	fmt.Println(v1)
	//  ["jom", "joh", "ja", "t"]
}
