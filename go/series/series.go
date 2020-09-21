package series

import (
	"strings"
)

// All returns a list of all substrings of s with length n.
func All(n int, s string) (res []string) {
	for i := range s {
		c := strings.Builder{}
		for j := 0; j < n; j++ {
			if i+j >= len(s) {
				break
			}
			c.WriteRune(rune(s[i+j]))
		}
		if c.Len() == n {
			res = append(res, c.String())
		}
	}
	return
}

// UnsafeFirst returns the first substring of s with length n.
func UnsafeFirst(n int, s string) string {

	c := strings.Builder{}
	for j := 0; j < n; j++ {
		if j >= len(s) {
			break
		}
		c.WriteRune(rune(s[j]))
	}
	if c.Len() != n {
		return ""
	}
	return c.String()
}
