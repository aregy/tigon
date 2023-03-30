package main

import "fmt"

type void struct{}

var member = void{}

func lengthOfLongestSubstring(s string) int {
	var length int

	for i := 0; i < len(s); i++ {
		x := uniqSequence(s[i:])
		if x > length {
			length = x
		}
	}
	return length
}

func uniqSequence(s string) int {
	uvals := make(map[int32]void)
	for i, v := range s {
		if _, ok := uvals[v]; ok {
			return i
		}
		uvals[v] = member
	}
	return len(s)
}

func display(s string) {
	fmt.Printf("The lonagest substring of unique characters within '%s' is of length %d\n", s, lengthOfLongestSubstring(s))
}

func main() {
	display("abcabcbb")
	display("bbbbb")
	display("pwwkew")
}
