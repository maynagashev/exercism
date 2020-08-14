/*
Package pangram contains functions for work with pangrams.

A pangram (Greek: παν γράμμα, pan gramma, "every letter") is a sentence using every letter of the alphabet at least once. The best known English pangram is:
The quick brown fox jumps over the lazy dog.

The alphabet used consists of ASCII letters a to z, inclusive, and is case insensitive. Input will not contain non-ASCII symbols.
*/
package pangram

import (
	"strings"
)

// IsPangram determines if a sentence is a pangram.
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	for i := byte('a'); i <= byte('z'); i++ {
		if strings.IndexByte(s, i) == -1 {
			return false
		}
	}
	return true
}
