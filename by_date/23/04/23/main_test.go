package uniq

import (
	"fmt"
	"testing"
)

func TestFindUniq(t *testing.T) {
	arr := []int{7, 3, 5, 5, 4, 3, 4, 8, 8}
	want := 7
	got := FindUniq(arr)
	if got != want {
		panic(fmt.Sprintf("%d in %v", got, arr))
	}
}
