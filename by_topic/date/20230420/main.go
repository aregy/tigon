package main

import "fmt"

type Pos struct {
	X, Y int
}

func (p Pos) Add(a, b int) Pos {
	p.X += a
	p.Y += b
	return p
}

func ProbBoardedPosition(x, y, k int) float64 {
	prev := make(map[Pos]int)
	prev[Pos{x, y}] = 1
	for i := 0; i < k; i++ {
		next := make(map[Pos]int)
		for p0 := range prev {
			next[p0.Add(1, 2)]++
			next[p0.Add(1, -2)]++
			next[p0.Add(-1, 2)]++
			next[p0.Add(-1, -2)]++
			next[p0.Add(2, 1)]++
			next[p0.Add(2, -1)]++
			next[p0.Add(-2, 1)]++
			next[p0.Add(-2, -1)]++
		}
		prev = next
	}
	var in int
	var out int
	fmt.Println(prev)
	for p, count := range prev {
		if 0 <= p.X && p.X <= 7 && 0 <= p.Y && p.Y <= 7 {
			in += count
		} else {
			out += count
		}
	}
	return float64(in) / (float64(out) + float64(in))
}

func main() {
	var a, b, c int
	for {
		fmt.Print(": ")
		fmt.Scanf("%d %d %d", &a, &b, &c)
		fmt.Printf("> %g\n", ProbBoardedPosition(a, b, c))
	}
}
