/*
Package wordcount contains function for counting words in phrases.

Rules:
1. The count is case insensitive (ie "You", "you", and "YOU" are 3 uses of the same word)
2. The count is unordered; the tests will ignore how words and counts are ordered
3. Other than the apostrophe in a contraction all forms of punctuation are ignored
4. The words can be separated by any form of whitespace (ie "\t", "\n", " ")

Benchmarks:
1. Regexp
BenchmarkWordCount-12              19478             62856 ns/op           43649 B/op        413 allocs/op
2. Current
BenchmarkWordCount-12             174574              5861 ns/op            4000 B/op         39 allocs/op
*/
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency return type for WordCount.
type Frequency = map[string]int

// WordCount counts the occurrences of each word in given string (unicode compatible).
func WordCount(s string) Frequency {
	res := Frequency{}
	words := strings.FieldsFunc(strings.ToLower(s), isSplitRune) // to lower & split
	for _, w := range words {
		w = strings.TrimFunc(w, isTrimRune) // trim
		res[w]++
	}
	return res
}

func isSplitRune(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '\''
}
func isTrimRune(r rune) bool {
	return !unicode.IsLetter(r) && !unicode.IsDigit(r)
}
