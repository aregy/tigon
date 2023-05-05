package main

import (
	"testing"
)

func TestCrop(t *testing.T) {
	src := "the quick brown fox jumps over the lazy dog"
	got := Crop(src, 10)
	want := []string{"the quick", "brown fox", "jumps over", "the lazy", "dog"}
	if !Comp(got, want) {
		t.Errorf("Got '%#v' in place of '%#v'", got, want)
	}
}
func Comp(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
