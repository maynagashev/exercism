package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

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

func (m Matrix) Cols() [][]int {
	return [][]int{}
}

func (m Matrix) Rows() [][]int {
	return [][]int{}
}

func (m Matrix) Set(r int, c int, val int) bool {
	return true
}

// validate checks if the "width" of all rows is the same
func (m Matrix) validate() error {
	width := len(m[0])
	for _, row := range m {
		if len(row) != width {
			return errors.New("invalid matrix")
		}
	}
	return nil
}