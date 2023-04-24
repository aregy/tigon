package main

import "fmt"

var CheckWS = map[rune]struct{}{
	rune(9):  struct{}{},
	rune(10): struct{}{},
	rune(32): struct{}{},
}

func ReverseWords(src string) string {
	res := []byte(src)
	bounds := make(map[int]int, len(src)/5)

	r := -1
	l := 0
	p := -1
	for i, letter := range src {
		if p != -1 && i-p > 1 {
			fmt.Printf(":%s\n", string(src[p:i+1]))
		}
		if r-l > 0 {
			bounds[l+1] = r
			r = l
		}
		_, curr := CheckWS[letter]
		if curr {
			l = i
		} else {
			r = i
		}
		p = i
	}
	for l, r := range bounds {
		for ; l <= r; l, r = l+1, r-1 {
			res[l], res[r] = res[r], res[l]
		}
	}
	for l, r := 0, len(res)-1; l <= r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return fmt.Sprintf("%s", res)
}

func main() {
	fmt.Print(ReverseWords(`can you read t®his®`))
}
