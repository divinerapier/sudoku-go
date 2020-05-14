package sudoku

import "testing"

func TestSudoku(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0, 2, 0, 0, 4, 8, 0},
		{5, 0, 0, 7, 0, 0, 0, 0, 0},
		{0, 0, 0, 5, 0, 8, 6, 2, 0},
		{0, 0, 5, 8, 7, 0, 0, 6, 1},
		{8, 0, 0, 0, 6, 0, 0, 0, 4},
		{6, 0, 0, 4, 3, 1, 9, 0, 8},
		{0, 9, 6, 1, 0, 4, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 4, 6},
		{0, 5, 4, 6, 0, 7, 0, 0, 0},
	}
	s, err := New(matrix)
	if err != nil {
		t.Fatalf("create sudoku. error: %v\n", err)
	}
	if s.colAvaliable(5, 0, 0) {
		t.Errorf("")
	}
}
