/*
Package atbash contains a method for encoding ascii string.

The Atbash cipher is a simple substitution cipher that relies on
transposing all the letters in the alphabet such that the resulting
alphabet is backwards. The first letter is replaced with the last
letter, the second with the second-last, and so on.

An Atbash cipher for the Latin alphabet would be as follows:

```text
Plain:  abcdefghijklmnopqrstuvwxyz
Cipher: zyxwvutsrqponmlkjihgfedcba
```

It is a very weak cipher because it only has one possible key, and it is
a simple monoalphabetic substitution cipher. However, this may not have
been an issue in the cipher's time.

Ciphertext is written out in groups of fixed length, the traditional group size
being 5 letters, and punctuation is excluded. This is to make it harder to guess
things based on word boundaries.

a-z 97-122

Benchmarks
BenchmarkAtbash-12        852914              1401 ns/op             688 B/op         29 allocs/op

*/
package atbash

import "strings"

// Atbash encodes given string with atbash cipher (a simple monoalphabetic substitution cipher)
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
	return string(b[:resLen])
}
