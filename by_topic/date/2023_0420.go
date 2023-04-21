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
	ensemble := make(map[int][]Pos)
	ensemble[0] = []Pos{Pos{x, y}}
	for i := 0; i < k; i++ {
		for _, p0 := range ensemble[i] {
			ensemble[i+1] = append(ensemble[i+1], p0.Add(1, 2))
			ensemble[i+1] = append(ensemble[i+1], p0.Add(1, -2))
			ensemble[i+1] = append(ensemble[i+1], p0.Add(-1, 2))
			ensemble[i+1] = append(ensemble[i+1], p0.Add(-1, -2))
			ensemble[i+1] = append(ensemble[i+1], p0.Add(2, 1))
			ensemble[i+1] = append(ensemble[i+1], p0.Add(2, -1))
			ensemble[i+1] = append(ensemble[i+1], p0.Add(-2, 1))
			ensemble[i+1] = append(ensemble[i+1], p0.Add(-2, -1))
		}
	}
	var in int
	var out int
	fmt.Println(ensemble[len(ensemble)-1])
	for _, p := range ensemble[len(ensemble)-1] {
		if 0 <= p.X && p.X <= 7 && 0 <= p.Y && p.Y <= 7 {
			in++
		} else {
			out++
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
