/*
Package cryptosquare contains the classic method for composing secret messages called a square code.

First, the input is normalized: the spaces and punctuation are removed from the English text and the message is downcased.
Then, the normalized characters are broken into rows. These rows can be regarded as forming a rectangle when printed with intervening newlines.
The plaintext should be organized in to a rectangle. The size of the rectangle (r x c) should be decided by the length of the message,
such that c >= r and c - r <= 1, where c is the number of columns and r is the number of rows.
*/
package main

import (
	"math"
	"strings"
	"unicode"
)

func main() {
	Normalize("First, the input is normalized: the spaces and punctuation are removed from the English text and the message is downcased.")
}


// Encode encodes string with square code algorithm.
func Encode(s string) string {
	var c, r int

	s = Normalize(s)
	if s == "" {
		return ""
	}
	sqrt := math.Sqrt(float64(len(s)))

	c = int(math.Round(sqrt))
	if sqrt > math.Round(sqrt) {
		r = c + 1
	} else {
		r = c
	}
	for i := r*c - len(s); i > 0; i-- {
		s = s + " "
	}
	builder := strings.Builder{}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			builder.WriteByte(s[i+(j*r)])
		}
		builder.WriteRune(' ')
	}
	s = builder.String()
	return s[:len(s)-1]
}
func Normalize(s string) string {
	s = strings.ToLower(s)
	builder := strings.Builder{}
	for _, el := range s {
		if unicode.IsDigit(el) || unicode.IsLetter(el) {
			builder.WriteRune(el)
		}
	}
	return builder.String()
}
