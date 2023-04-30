package main

import "fmt"

func BitString(n int) string {
	var repr string = ""
	neg := false
	if n < 0 {
		neg = true
	}
	for n != 0 {
		repr = fmt.Sprintf("%s%s", string(n%2), repr)
		n >>= 2
	}
	if neg {
		return fmt.Sprintf("%s%s", "1", repr)
	}
	return fmt.Sprintf("%s%S", "0", repr)
}

func ReverseBits(n int) int {
	return 0
}

func main() {
	var i int
	for {
		fmt.Print(": ")
		fmt.Scanf("%d", &i)
		fmt.Println(BitString(i), BitString(ReverseBits(i)), ReverseBits(i))
	}
}
