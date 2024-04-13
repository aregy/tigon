package main

import (
	"bufio"
	"fmt"
	"os"
)

func Reverse(str string) string {
	var arr []rune = []rune(str)
	for b, e := 0, len(arr)-1; b < e; b, e = b+1, e-1 {
		arr[b], arr[e] = arr[e], arr[b]
	}
	return string(arr)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		panic(err)
	}
	fmt.Print(Reverse(scanner.Text()))
}
