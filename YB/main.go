package main

import (
	"bytes"
	"fmt"
	"reflect"
	"time"
)

func Eq(a, b []any) bool {
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

func PrintEq(a, b any) {
	// if Eq(a, b) {
	if reflect.DeepEqual(a, b) {
		fmt.Printf("%v and %v are equal\n", a, b)
	} else {
		fmt.Printf("%v and %v are not equal\n", a, b)
	}
}

func main() {
	str1 := "Yabdabaocehuncaoheuncahoecnuhaocneuhchcnaeouaaaaaa"
	str2 := "Yabdabaocehuncaoheuncahoecnuhaocneuhchcnaeouaoceuh"
	t1 := time.Now()
	res := bytes.Equal([]byte(str1), []byte(str2))
	t2 := time.Now()
	dur1 := t2.Sub(t1)
	fmt.Println(dur1, res)

}
