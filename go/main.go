package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(Abbreviate("    sdf sdfd.qwe qwe-12q"))
}

func isSplitRune(c rune) bool {
	return !unicode.IsLetter(c) && c != '\''
}

// Abbreviate splits given string to words by whitespaces (defined by unicode.IsSpace) AND by hyphen,
// then joins all first letters in words
func Abbreviate(s string) (res string) {
	words := strings.FieldsFunc(s, isSplitRune) // split string
	fmt.Println(words)
	for _, w := range words {
		r := []rune(w) // convert to runes for proper unicode handling
		if len(r) > 0 {
			res += string(r[0])
		}
	}
	return strings.ToUpper(res)
}

func isNotLetter(c rune) bool {
	return !unicode.IsLetter(c)
}
