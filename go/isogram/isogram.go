/*
Package isogram containts method for determining if word is isogram.

An isogram (also known as a "nonpattern word") is a word or phrase without a repeating letter,
however spaces and hyphens are allowed to appear multiple times.
*/
package isogram

import "unicode"

// IsIsogram determines if word is isogram
func IsIsogram(s string) bool {
	seen := map[rune]bool{}
	for _, r := range s {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		if seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}
