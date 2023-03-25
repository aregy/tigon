package main

import "fmt"

type Matrix [][]any

func (matrix Matrix) String() string {
	var s string
	for m, row := range matrix {
		for n, _ := range row {
			s += fmt.Sprintf(" %v ", matrix[m][n])
		}
		s += "\n"
	}
	return s
}
func CreateNewMatrix(m, n int) *Matrix {
	var matrix Matrix = make([][]any, m)
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]any, n)
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] = 0
		}
	}
	return &matrix
}

func InitializeMatrix(rows ...[]any) *Matrix {
	if len(rows) == 0 || len(rows[0]) == 0 {
		panic("Cannot initialize matrices with a 0 dimension")
	}
	width := len(rows[0])
	for i, _ := range rows {
		if len(rows[i]) != width {
			panic("Cannot initialize a matrix with a ragged array")
		}
	}
	matrix := CreateNewMatrix(len(rows), len(rows[0]))
	if len(rows) > len(*matrix) {
		panic(fmt.Sprintf("len(rows) <%d> > len(*matrix) <%d>", len(rows), len(*matrix)))
	}
	for m, row := range rows {
		for n, val := range row {
			(*matrix)[m][n] = val
		}
		//copy(row, (*matrix)[m])
	}
	return matrix
}
func (matrix *Matrix) Rotate() {
	rotated := CreateNewMatrix(len((*matrix)[0]), len((*matrix)))
	for m, row := range *rotated {
		for n, _ := range row {
			(*rotated)[m][n] = (*matrix)[n][m]
		}
	}
	*matrix = *rotated
}

func main() {
	mat := InitializeMatrix([]any{1, 2, 3, 4}, []any{5, 6, 7, 8})

	fmt.Println(mat)
	mat.Rotate()
	fmt.Println(mat)
}
