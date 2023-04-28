package main

import (
	"bufio"
	"fmt"
	"os"
)

type F struct {
	Letter rune
	Count  int
}
type P struct {
	Entry  string
	Weight int
}

func (f F) String() string {
	return fmt.Sprintf("%s:%d", string(f.Letter), f.Count)
}

func qsort_fn1(fs []F, conv func(any, int) int) {
	var qs func(int, int)
	partition := func(l, r int) int {
		e := conv(fs, r)
		k := l - 1
		for i := l; i < r; i++ {
			if conv(fs, i) >= e {
				k++
				fs[i], fs[k] = fs[k], fs[i]
			}
		}
		fs[k+1], fs[r] = fs[r], fs[k+1]
		return k + 1
	}
	qs = func(low, hi int) {
		if low < hi {
			pivot := partition(low, hi)
			qs(low, pivot-1)
			qs(pivot+1, hi)
		}
	}
	qs(0, len(fs)-1)
}
func qsort_fn2(fs []P, conv func(any, int) int) {
	var qs func(int, int)
	partition := func(l, r int) int {
		e := conv(fs, r)
		k := l - 1
		for i := l; i < r; i++ {
			if conv(fs, i) >= e {
				k++
				fs[i], fs[k] = fs[k], fs[i]
			}
		}
		fs[k+1], fs[r] = fs[r], fs[k+1]
		return k + 1
	}
	qs = func(low, hi int) {
		if low < hi {
			pivot := partition(low, hi)
			qs(low, pivot-1)
			qs(pivot+1, hi)
		}
	}
	qs(0, len(fs)-1)
}

func main() {
	file, e := os.Open("./src.txt")
	if e != nil {
		panic(e)
	}
	defer file.Close()
	//var b []byte
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	catalog := make(map[rune]int)
	lines := 0
	values := make(map[string]struct{})

	rejects := []string{}
	for scanner.Scan() {
		if len(scanner.Text()) != 10 {
			rejects = append(rejects, scanner.Text())
			continue
		}

		for _, k := range scanner.Text() {
			//fmt.Print(k, fmt.Sprintf("(%s)|", string(k)))
			catalog[k]++
		}
		values[scanner.Text()] = struct{}{}
		lines++
		//	fmt.Println(": ", scanner.Text(0))

	}
	pairs := make([]P, len(values))
	fmt.Printf("%d entries processed\n", lines)
	freq := make([]F, len(catalog))
	i := 0
	for l, v := range catalog {
		freq[i] = F{l, v}
		i++
	}
	fmt.Printf("%v\n", freq)
	qsort_fn1(freq, func(src any, i int) int {
		list, ok := src.([]F)
		if !ok {
			panic(list)
		}
		return list[i].Count
	})
	fmt.Printf("%d items in freq\n", len(freq))
	fmt.Printf("%v\n", freq)
	i = -1
	for entry, _ := range values {
		i++
		pairs[i] = P{}
		pairs[i].Entry = entry
		letterValues := make(map[int]struct{})
		for _, letter := range entry {
			letterValues[int(letter)] = struct{}{}
		}
		for lv, _ := range letterValues {
			pairs[i].Weight += lv
		}
	}
	qsort_fn2(pairs, func(src any, i int) int {
		list, ok := src.([]P)
		if !ok {
			panic(list)
		}
		return list[i].Weight
	})
	used := make(map[rune]struct{})
	k := -1
	for _, pair := range pairs {
		valid := true
		for _, letter := range pair.Entry {
			if _, ok := used[letter]; ok {
				valid = false
			}
			used[letter] = struct{}{}
		}
		if !valid {
			continue
		}
		k++
		fmt.Printf("#%d. %s (%d)\n", k+1, pairs[i].Entry, pairs[i].Weight)
	}
	fmt.Printf("with %d rejects not in %d: %v\n", len(rejects), len(pairs), rejects)
}
