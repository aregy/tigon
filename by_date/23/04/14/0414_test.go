package main

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	phrase := `can you read this`

	got := ReverseWords(phrase)
	want := `this read you can`
	if got != want {
		panic(fmt.Sprintf("'%q' in place of '%q'", got, want))
	}
}
