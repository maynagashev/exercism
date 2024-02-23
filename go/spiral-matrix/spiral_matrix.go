package spiralmatrix

// SpiralMatrix generates a square matrix filled with numbers in a spiral order.
func SpiralMatrix(size int) [][]int {
	// Initialize the matrix with the given size.
	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i] = make([]int, size)
	}

	// Initialize boundaries and the starting number.
	left, top, right, bottom := 0, 0, size-1, size-1
	num := 1

	// Fill the matrix in a spiral order.
	for left <= right && top <= bottom {
		// Fill the top row, moving right.
		for i := left; i <= right; i++ {
			matrix[top][i] = num
			num++
		}
		top++

		// Fill the right column, moving down.
		for i := top; i <= bottom; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// Fill the bottom row, moving left.
		for i := right; i >= left; i-- {
			matrix[bottom][i] = num
			num++
		}
		bottom--

		// Fill the left column, moving up.
		for i := bottom; i >= top; i-- {
			matrix[i][left] = num
			num++
		}
		left++
	}

	return matrix
}
