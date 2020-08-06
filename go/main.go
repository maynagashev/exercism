package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "  foo bar  baz   The Road _Not_ Taken Halley's Comet Complementary metal-oxide semiconductor Something - I made up from thin air"
	s = strings.ReplaceAll(s, "-", "- ")
	words := strings.Fields(s)
	res := ""

	for _, w := range words {
		w = strings.TrimFunc(w, func(c rune) bool {
			return !unicode.IsLetter(c)
		})
		if w!="" {
			res += strings.ToUpper(string(w[0]))
		}
	}

	fmt.Printf("Fields are: %q \n %s", strings.Fields(s), res)
}
