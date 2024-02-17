package diamond

import (
	"errors"
	"strings"
)

// Gen generates a diamond shape string starting with 'A' up to the input char.
func Gen(char byte) (string, error) {
	// Check if input is within 'A' to 'Z'
	if char < 'A' || char > 'Z' {
		return "", errors.New("input must be a capital letter A-Z")
	}
	fill := ' '

	// Determine the size of the diamond based on the input char
	size := int(char - 'A')

	// Initialize a builder to construct the diamond string
	var builder strings.Builder

	// Generate the diamond
	for i := 0; i <= size*2; i++ {
		// Calculate current row's letter (mirror in the bottom half)
		// size: 2, row: 0 1 2 1 0
		row := i
		if i > size {
			row = size*2 - i
		}

		// Leading spaces
		for j := 0; j < size-row; j++ {
			builder.WriteRune(fill)
		}

		// Current letter based on row
		builder.WriteByte('A' + byte(row))

		// Spaces between letters, if not the first row
		if row > 0 {
			for j := 0; j < 2*row-1; j++ {
				builder.WriteRune(fill)
			}
			builder.WriteByte('A' + byte(row))
		}

		// Trailing spaces
		for j := 0; j < size-row; j++ {
			builder.WriteRune(fill)
		}

		// Add a newline if not the last line
		if i < size*2 {
			builder.WriteRune('\n')
		}
	}

	return builder.String(), nil
}
