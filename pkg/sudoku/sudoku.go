package sudoku

import (
	"errors"
	"fmt"
)

type Sudoku struct {
	matrix [][]int
}

func New(matrix [][]int) (*Sudoku, error) {
	if len(matrix) != 9 {
		return nil, errors.New("should be 9 rows")
	}

	for i := range matrix {
		if len(matrix[i]) != 9 {
			return nil, errors.New("should be 9 rows")
		}
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < 0 || matrix[i][j] > 9 {
				return nil, errors.New("value should be in range [0, 9]")
			}
		}
	}

	return &Sudoku{
		matrix: matrix,
	}, nil
}

func (s *Sudoku) Compute() (*Sudoku, error) {
	copied := s.copy()
	if copied.dfs(0, 0) {
		return copied, nil
	}
	return nil, errors.New("no answer")
}

func (s *Sudoku) copy() *Sudoku {
	results := make([][]int, 9)
	for i := 0; i < 9; i++ {
		results[i] = make([]int, 9)
		for j := 0; j < 9; j++ {
			results[i][j] = s.matrix[i][j]
		}
	}
	return &Sudoku{matrix: results}
}

func (s *Sudoku) rowAvaliable(v, i, j int) bool {
	for n := 0; n < 9; n++ {
		if n == j {
			continue
		}
		if v == s.matrix[i][n] {
			return false
		}
	}
	return true
}

func (s *Sudoku) colAvaliable(v, i, j int) bool {
	for n := 0; n < 9; n++ {
		if n == i {
			continue
		}
		if v == s.matrix[n][j] {
			return false
		}
	}
	return true
}

func (s *Sudoku) squareAvaliable(v, i, j int) bool {
	r := i % 3
	c := j % 3
	for m := 0; m < 3; m++ {
		for n := 0; n < 3; n++ {
			if m == r && n == c {
				continue
			}
			x := i/3*3 + m
			y := j/3*3 + n
			if s.matrix[x][y] == v {
				return false
			}
		}
	}
	return true
}

func (s *Sudoku) dfs(i, j int) bool {
	if i == 9 {
		return true
	}
	if s.matrix[i][j] != 0 {
		return s.dfs(i+(j+1)/9, (j+1)%9)
	}

	for v := 1; v <= 9; v++ {
		if !s.isAvailable(v, i, j) {
			continue
		}
		s.matrix[i][j] = v
		if s.dfs(i+(j+1)/9, (j+1)%9) {
			return true
		}
	}
	s.matrix[i][j] = 0
	return false
}

func (s *Sudoku) Dump() string {
	var str string
	for i := 0; i < 9; i++ {
		line := ""
		for j := 0; j < 9; j++ {
			line += fmt.Sprintf(" %d", s.matrix[i][j])
		}
		str = str + line + "\n"
	}
	return str
}

func (s *Sudoku) isAvailable(v, i, j int) bool {
	return s.colAvaliable(v, i, j) &&
		s.rowAvaliable(v, i, j) &&
		s.squareAvaliable(v, i, j)
}
