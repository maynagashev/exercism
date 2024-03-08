package railfence

func Encode(message string, rails int) string {
	matrix := make([][]byte, rails)
	for i := range matrix {
		matrix[i] = make([]byte, len(message)) // Initialize with spaces
	}

	row, dir := 0, 1
	for i, char := range message {
		matrix[row][i] = byte(char)
		row += dir
		// Change direction when reaching the top or bottom rail
		if row == rails-1 || row == 0 {
			dir = -dir
		}
	}

	// Read characters from the matrix row by row, removing spaces
	encoded := make([]byte, 0)
	for _, row := range matrix {
		for _, char := range row {
			if char != '\x00' { // Skip padding spaces
				encoded = append(encoded, char)
			}
		}
	}

	return string(encoded)
}

// Decode decodes a message encoded with the rail fence cipher
// with the given number of rails.
// It creates a matrix with the same dimensions as the encoded message,
// then fills it with the message characters in the same pattern as the encoding.

func Decode(message string, rails int) string {
	matrix := make([][]byte, rails)
	for i := range matrix {
		matrix[i] = make([]byte, len(message)) // Initialize with spaces
	}

	// Fill the matrix with placeholders where the characters will be placed
	row, dir := 0, 1
	for i, _ := range message {
		matrix[row][i] = 1
		row += dir
		// Change direction when reaching the top or bottom rail
		if row == rails-1 || row == 0 {
			dir = -dir
		}
	}

	// Fill the matrix with the characters from the message
	index := 0
	for i, rowChars := range matrix {
		for j, char := range rowChars {
			if char == 1 { // Skip padding spaces
				matrix[i][j] = message[index]
				index++
			}
		}
	}

	// Read characters from the matrix as if it was encoded
	decoded := make([]byte, 0)
	row, dir = 0, 1
	for i := 0; i < len(message); i++ {
		decoded = append(decoded, matrix[row][i])
		row += dir
		// Change direction when reaching the top or bottom rail
		if row == rails-1 || row == 0 {
			dir = -dir
		}
	}

	return string(decoded)
}
