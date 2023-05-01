package lib

import (
	"fmt"
	"strconv"
)

const origin = 1329 - 1 // immediately preceding the Ա rune

func Base36Digits(x int) string {
	var numerals string
	for x > 0 {
		if r := x % 36; r == 0 {
			numerals += "0"
		} else {
			numerals = string(origin+r) + numerals
		}
		x /= 36
	}
	return numerals
}

func traditionalNumerals(x int) string {
	if x == 0 {
		return "NaN"
	}
	var numerals string
	decades := 0
	for x > 0 {
		r := x % 10
		if r != 0 {
			numerals = string(origin+decades+r) + numerals
		}
		x /= 10
		decades++
	}
	return numerals
}

func TestBase36Digits() {
	for {
		var uInput string
		fmt.Print(": ")
		fmt.Scanln(&uInput)
		if x, err := strconv.Atoi(uInput); err != nil {
			fmt.Errorf("%s is not integer-convertible\n", uInput)
		} else {
			fmt.Printf("%s in traditional\n", traditionalNumerals(x))
			fmt.Printf("%s (base-36)\n", Base36Digits(x))
		}
		fmt.Println()
	}
}
