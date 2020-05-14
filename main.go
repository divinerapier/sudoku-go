package main

import (
	"fmt"

	"github.com/divinerapier/sudoku-go/pkg/sudoku"
)

func main() {
	matrix := [][]int{
		{6, 0, 0, 0, 0, 0, 9, 1, 5},
		{0, 0, 0, 0, 0, 7, 0, 0, 0},
		{1, 0, 0, 0, 5, 0, 0, 0, 0},
		{3, 0, 2, 0, 0, 8, 0, 0, 0},
		{0, 6, 0, 5, 0, 1, 0, 3, 0},
		{0, 0, 0, 4, 0, 0, 5, 0, 9},
		{0, 0, 0, 0, 8, 0, 0, 0, 4},
		{0, 0, 0, 6, 0, 0, 0, 0, 0},
		{4, 7, 3, 0, 0, 0, 0, 0, 2},
	}
	s, err := sudoku.New(matrix)
	if err != nil {
		panic(err)
	}
	result, err := s.Compute()
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Dump())
}
