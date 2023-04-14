package main

import "fmt"

var powers map[uint]struct{}

func init() {
	powers = make(map[uint]struct{})
	for d := 2; d < 32; d += 2 {
		powers[uint(1)<<d] = struct{}{}
	}
}

func isFourPower(n uint) bool {
	if _, exists := powers[n]; exists {
		return true
	}
	return false
}

func stat(n uint) {
	if isFourPower(n) {
		fmt.Printf("%d is a power of four\n", n)
	} else {
		fmt.Printf("%d is not a power of four\n", n)
	}

}

func main() {
	stat(uint(3))
	stat(uint(4))
	stat(uint(64))
}
