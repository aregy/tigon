package main

import "fmt"

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n),
	}
	for idx, _ := range uf.parent {
		uf.parent[idx] = idx + 1
	}
	return uf
}
func (uf *UnionFind) Find(x int) int {
	for true { 
		if uf.parent[x - 1] == x {
		    break
		}
		x = uf.parent[x - 1]
	}
	return x
}

func (uf *UnionFind) Union(x, y int) bool {
	fx, fy := uf.Find(x), uf.Find(y)
	if fx == fy {
		return false
	}
	m := max(fx, fy)
	uf.parent[fx - 1], uf.parent[fy - 1] = m, m 
	return true
}

func (uf *UnionFind) GetCycle(singly_connected_tree [][]int) []int {
	for _, edge := range singly_connected_tree {
		if !uf.Union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}

func main() {
	testcases := [][][]int{{{1, 2}, {1, 3}, {2, 3}}}
	testcases = append(testcases, [][]int{{1,2},{2,3},{3,4},{1,4},{1,5}})
	for _, tc := range testcases {
		fmt.Printf("Given: %v\n", tc)
		uf := NewUnionFind(len(tc))
		if edge := uf.GetCycle(tc); edge != nil {
			fmt.Printf("Cycle at %v\n", edge)
		} else {
			fmt.Println("No cycle present!")
		}
	}
}
