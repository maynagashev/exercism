package acronym

import (
	"strings"
	"unicode"
)

func isSplitRune(c rune) bool {
	return !unicode.IsLetter(c) && c != '\''
}

// Abbreviate splits given string to words then joins all first letters in words
func Abbreviate(s string) (res string) {
	words := strings.FieldsFunc(s, isSplitRune) // split string
	for _, w := range words {
		r := []rune(w) // convert to runes for proper unicode handling
		if len(r) > 0 {
			res += string(r[0])
		}
	}
	return strings.ToUpper(res)
}
