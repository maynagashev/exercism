/*
Package cryptosquare contains the classic method for composing secret messages called a square code.

First, the input is normalized: the spaces and punctuation are removed from the English text and the message is downcased.
Then, the normalized characters are broken into rows. These rows can be regarded as forming a rectangle when printed with intervening newlines.
The plaintext should be organized in to a rectangle. The size of the rectangle (r x c) should be decided by the length of the message,
such that c >= r and c - r <= 1, where c is the number of columns and r is the number of rows.
*/
package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Encode("ZOMG! ZOMBIES!!!"))
}

// Encode encodes string with square code algorithm.
func Encode(s string) (res string) {
	var c, r int
	s = normalize(s)
	if s == "" {
		return ""
	}
	r, c = rectDims(len(s))
	fmt.Println(s, len(s), c, r)

	// add trailing space to input string?
	for i := r*c - len(s); i > 0; i-- {
		s = s + " "
	}

	builder := strings.Builder{}
	for resRow := 0; resRow < c; resRow++ {
		for resCol := 0; resCol < r; resCol++ {
			builder.WriteByte(s[resRow+resCol*c])
		}
		builder.WriteRune(' ')
	}
	res = builder.String()

	return res
}

func normalize(s string) string {
	if s == "" {
		return ""
	}
	builder := strings.Builder{}
	builder.Grow(len(s))
	for _, el := range s {
		if unicode.IsDigit(el) || unicode.IsLetter(el) {
			builder.WriteRune(unicode.ToLower(el))
		}
	}
	return builder.String()
}

func rectDims(l int) (r, c int) {
	sqrt := math.Sqrt(float64(l))
	roundSqrt := math.Round(sqrt)
	r = int(roundSqrt)
	c = int(roundSqrt)
	if  sqrt > roundSqrt {
		c = r + 1
	}
	return
}