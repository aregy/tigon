package main

import (
	"fmt"
	"strings"
)

func BitString(n int) string {
	repr := []string{}
	isNeg := false
	pos, neg := func() (string, string) { return "1", "0" }()
	if n < 0 {
		isNeg = true
		n *= -1
		n -= 1
		pos = "0"
		neg = "1"
	}
	for n != 0 {
		if n%2 == 0 {
			repr = append([]string{neg}, repr...)
		} else {
			repr = append([]string{pos}, repr...)
		}
		n >>= 1
	}
	if isNeg {
		repr = append([]string{"1"}, repr...)
	} else {
		repr = append([]string{"0"}, repr...)
	}
	return strings.Join(repr, "")
}

// TODO
func ReverseBits(n int) int {
	u := uint(n)
	res := 0
	if n < 0 {
		res += 1
	}
	res <<= 1
	for u > 0 {
		if res%2 == 0 {
			res += 1
		}
		res <<= 1
		u >>= 1
	}
	return int(res)
}

func main() {
	var i int
	for {
		fmt.Print(": ")
		fmt.Scanf("%d", &i)
		fmt.Println(BitString(i), BitString(ReverseBits(i)), ReverseBits(i))
	}
}
