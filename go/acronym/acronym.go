package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate splits given string to words by whitespaces (defined by unicode.IsSpace) AND by hyphen,
// then joins all first letters in words
func Abbreviate(s string) (res string) {
	words := strings.Fields(strings.ReplaceAll(s, "-", " ")) // split string by whitespaces AND hyphen

	for _, w := range words {
		w = strings.TrimLeftFunc(w, isNotLetter) // trimming all non letters from left
		r := []rune(w)                           // convert to runes for proper unicode handling
		if len(r) > 0 {
			res += string(r[0])
		}
	}
	return strings.ToUpper(res)
}

func isNotLetter(c rune) bool {
	return !unicode.IsLetter(c)
}
