package main

import (
	"testing"
)

func TestBitString(t *testing.T) {
	want := "0000 0000 0000 0000 0000 0000 0010 0001"
	got := BitString(33)
	if want != got {
		t.Errorf("got '%v' but expected '%v'", got, want)
	}
}
func TestReverseBits(t *testing.T) {
	in := 33
	want := 0

	for _, v := range BitString(in) {
		if string(v) == "1" {
			want += 1
		}
		want *= 2
	} //"1000 0100 0000 0000"
	got := ReverseBits(in)
	if want != got {
		t.Errorf("got '%v' but expected '%v'", got, want)
	}
}
