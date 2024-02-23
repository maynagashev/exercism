package wordsearch

import "errors"

// Directions to search: horizontal, vertical, and two diagonals
var directions = [][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0}, // right, down, left, up
	{1, 1}, {-1, -1}, {1, -1}, {-1, 1}, // diagonal: down-right, up-left, down-left, up-right
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	result := make(map[string][2][2]int)
	for _, word := range words {
		found := false
		for rowIndex, row := range puzzle {
			for colIndex := range row {
				for _, dir := range directions {
					if checkWord(word, rowIndex, colIndex, dir, puzzle) {
						result[word] = [2][2]int{{colIndex, rowIndex}, {colIndex + dir[1]*(len(word)-1), rowIndex + dir[0]*(len(word)-1)}}
						found = true
						break
					}
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			// If word not found, return error as specified in test cases
			return nil, errors.New("word not found")
		}
	}
	return result, nil
}

// checkWord checks if a word exists in the given direction from the start position
func checkWord(word string, row, col int, dir [2]int, puzzle []string) bool {
	for i := 0; i < len(word); i++ {
		r := row + i*dir[0]
		c := col + i*dir[1]
		// Check if the word is out of bounds or the character does not match
		if r < 0 || r >= len(puzzle) || c < 0 || c >= len(puzzle[0]) || puzzle[r][c] != word[i] {
			return false
		}
	}
	return true
}
