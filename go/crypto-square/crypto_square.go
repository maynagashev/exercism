/*
Package cryptosquare contains the classic method for composing secret messages called a square code.

First, the input is normalized: the spaces and punctuation are removed from the English text and the message is downcased.
Then, the normalized characters are broken into rows. These rows can be regarded as forming a rectangle when printed with intervening newlines.
*/
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode encodes string with square code algorithm.
func Encode(s string) string {
	if s == "" {
		return ""
	}

	s = normalize(s)
	r, c := rectDims(len(s))

	builder := strings.Builder{}
	builder.Grow(len(s))
	// resulting rows count = count of input rectangle cols (c)
	for resRow := 0; resRow < c; resRow++ {
		// resulting cols count = count of input rectangle rows (r)
		for resCol := 0; resCol < r; resCol++ {
			// pick every c-th symbol (resCol*c) in current resRow
			i := resCol*c + resRow
			// if current index out of bounds (string ended), then add trailing space in resulting row
			if i >= len(s) {
				builder.WriteByte(' ')
				continue
			}
			builder.WriteByte(s[i])
		}

		// output chunks (rows) separator (space symbol), without trailing space
		if resRow < c-1 {
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}

// normalize removes all non-letter and non-digit runes from given string
func normalize(s string) string {
	builder := strings.Builder{}
	builder.Grow(len(s))
	for _, el := range s {
		if unicode.IsDigit(el) || unicode.IsLetter(el) {
			builder.WriteRune(unicode.ToLower(el))
		}
	}
	return builder.String()
}

// rectDims calculates dimensions for source rectangle
// The size of the rectangle (r x c) should be decided by the length of the message,
// such that c >= r and c - r <= 1, where c is the number of columns and r is the number of rows.
func rectDims(l int) (r, c int) {
	sqrt := math.Sqrt(float64(l))
	roundSqrt := math.Round(sqrt)
	r = int(roundSqrt)
	c = int(roundSqrt)
	if sqrt > roundSqrt {
		c = r + 1
	}
	return
}
