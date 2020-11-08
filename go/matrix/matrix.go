package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix representation
type Matrix [][]int

// New constructs new matrix from string input
func New(input string) (Matrix, error) {
	rows := strings.Split(input, "\n")
	m := make(Matrix, len(rows))
	for i, r := range rows {
		cols := strings.Fields(r)
		m[i] = make([]int, len(cols))
		for j, c := range cols {
			v, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}
			m[i][j] = v
		}
	}
	return m, m.validate()
}

// Cols makes copy of matrix and rotates it
func (m Matrix) Cols() [][]int {
	cols := make([][]int, len(m[0]))
	for i := 0; i < len(cols); i++ {
		cols[i] = make([]int, len(m))
		for j := 0; j < len(cols[i]); j++ {
			cols[i][j] = m[j][i]
		}
	}
	return cols
}

// Rows makes copy of the matrix and returns as is
func (m Matrix) Rows() [][]int {
	rows := make([][]int, len(m))
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]int, len(m[i]))
		for j := 0; j < len(rows[i]); j++ {
			rows[i][j] = m[i][j]
		}
	}
	return rows
}

// Set updates value in specified cell
func (m Matrix) Set(r int, c int, val int) (ok bool) {
	if r < 0 || c < 0 || r >= len(m) || c >= len(m[0]) {
		return false
	}
	m[r][c] = val
	return true
}

// validate checks if the "width" of all rows in the matrix is the same
func (m Matrix) validate() error {
	width := len(m[0])
	for _, row := range m {
		if len(row) != width {
			return errors.New("invalid matrix")
		}
	}
	return nil
}
