package main

import (
	"fmt"
	"strconv"
)

func BinAdd(a, b string) string {
	num1, _ := strconv.ParseInt(a, 2, 64)
	num2, _ := strconv.ParseInt(b, 2, 64)
	return strconv.FormatInt(num1+num2, 2)
}

func main() {
	fmt.Println(BinAdd("11101", "1011"))
}
