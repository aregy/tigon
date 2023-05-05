package main

import "fmt"

// This problem was asked by Amazon.
//
// Given a string s and an integer k, break up the string into multiple lines such that each line has a length of k or less. You must break it up so that words don't break across lines. Each line has to have the maximum possible amount of words. If there's no way to break the text up, then return null.
//
// You can assume that there are no spaces at the ends of the string and that there is exactly one space between each word.
func Crop(src string, width int) []string {
	for _, v := range src {
		if len(v) > width {
			return nil
		}
	}
	delimiters := map[string]struct{} { " ": struct{}{}}
	prevLine := 0
	word := 0
	prevWord := 0
	for i, letter := range src {

		if _, ok := delimiters[letter]; ok && i > 1 {
			prevWord = i - 1
			word = i -1 - prevLine
		}
	}
	res := []string{""}
	c := 0
	(len(res[len(res) -1])+)
	return []string{}
}

func main() {
	src := "the quick brown fox jumps over the lazy dog"
	fmt.Println(Crop(src, 10))
}
