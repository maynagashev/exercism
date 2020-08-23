/*
a-z 97-122
*/
package atbash

import "strings"

func Atbash(s string) string {
	s = strings.ToLower(s)
	b := make([]byte, len(s)+len(s)/5)
	resLen, spaces := 0, 0
	for i := 0; i < len(s); i++ {

		// skip non letters and non digits
		if s[i] < '0' || (s[i] > '9' && s[i] < 'a') || s[i] > 'z' {
			continue
		}

		// spacing in output
		if spaces >= 5 {
			b[resLen] = ' '
			resLen++
			spaces = 0
		}

		// digits passing as is
		if s[i] >= '0' && s[i] <= '9' {
			b[resLen] = s[i]
		} else {
			b[resLen] = 'z' - s[i] + 'a'
		}
		resLen++
		spaces++

	}
	return string(b)
}
