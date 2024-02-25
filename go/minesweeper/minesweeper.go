package minesweeper

// Annotate returns an annotated board.
func Annotate(board []string) []string {
	rowCount := len(board)
	if rowCount == 0 {
		return board
	}
	colCount := len(board[0])

	// Convert the board to a mutable slice of byte slices.
	mutableBoard := make([][]byte, rowCount)
	for i := range board {
		mutableBoard[i] = []byte(board[i])
	}

	// Define a function to count adjacent mines.
	countMines := func(row, col int) int {
		count := 0
		for dx := -1; dx <= 1; dx++ {
			for dy := -1; dy <= 1; dy++ {
				checkRow, checkCol := row+dx, col+dy
				// Check if the new row and column are within the bounds of the board.
				if checkRow >= 0 && checkRow < rowCount && checkCol >= 0 && checkCol < colCount {
					if mutableBoard[checkRow][checkCol] == '*' {
						count++
					}
				}
			}
		}
		return count
	}

	// Annotate the board.
	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			if mutableBoard[row][col] != '*' {
				// Count the number of adjacent mines.
				mines := countMines(row, col)
				if mines > 0 {
					// If there are adjacent mines, annotate the cell with the number of adjacent mines.
					mutableBoard[row][col] = byte('0' + mines)
				}
			}
		}
	}

	// Convert the mutable board back to a slice of strings.
	annotatedBoard := make([]string, rowCount)
	for i, row := range mutableBoard {
		annotatedBoard[i] = string(row)
	}

	return annotatedBoard
}
