package main

import (
	"fmt"
	"strconv"
	"strings"
)

func bitRepr(x int) string {
	var s string = ""
	var signedness string = "0"
	if x == 0 {
		return "0"
	} else if x < 0 {
		signedness = "1"
		x *= -1
	}
	for x != 0 {
		r := x % 2
		s = strings.Join([]string{strconv.Itoa(r), s}, "")
		x /= 2
	}
	return signedness + s
}
func main() {
	fmt.Println(bitRepr(255))
	fmt.Println(bitRepr(-255))
	fmt.Println(bitRepr((^0 >> 1)))
}
