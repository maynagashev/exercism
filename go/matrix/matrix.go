package matrix

type Matrix string

func New(m string) (Matrix, error) {

	return Matrix(""), nil
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
