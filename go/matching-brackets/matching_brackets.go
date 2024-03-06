package brackets

func Bracket(input string) bool {
	var stack []rune

	// Mapping of closing brackets to their corresponding opening brackets
	matchingBrackets := map[rune]rune{']': '[', '}': '{', ')': '('}

	for _, char := range input {
		switch char {
		case '[', '{', '(':
			// Push opening brackets onto the stack
			stack = append(stack, char)
		case ']', '}', ')':
			// Check for matching opening brackets when a closing bracket is found
			if len(stack) == 0 || stack[len(stack)-1] != matchingBrackets[char] {
				// No matching opening bracket or stack is empty
				return false
			}
			// Pop the last opening bracket from the stack
			stack = stack[:len(stack)-1]
			// Ignore all other characters
		}
	}

	// If the stack is empty, all brackets were properly closed and nested
	return len(stack) == 0
}
