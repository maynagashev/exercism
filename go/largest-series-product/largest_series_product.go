/*
Package lsproduct (Largest Series Product)

Given a string of digits, calculate the largest product for a contiguous
substring of digits of length n.

For example, for the input `'1027839564'`, the largest product for a
series of 3 digits is 270 (9 * 5 * 6), and the largest product for a
series of 5 digits is 7560 (7 * 8 * 3 * 9 * 5).

Note that these series are only required to occupy *adjacent positions*
in the input; the digits need not be *numerically consecutive*.

For the input `'73167176531330624919225119674426574742355349194934'`,
the largest product for a series of 6 digits is 23520.

Benchmarks:

1. strconv.Atoi(string(byte))
BenchmarkLargestSeriesProduct-12          360379              3270 ns/op              64 B/op          4 allocs/op

2. hacking with digit bytes (b - '0')
BenchmarkLargestSeriesProduct-12         2005789               616 ns/op              64 B/op          4 allocs/op

*/
package lsproduct

import (
	"errors"
)

//LargestSeriesProduct calculates the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(s string, span int) (maxP int64, err error) {
	if span > len(s) || span < 0 {
		return 0, errors.New("span is out of bounds")
	}
	for i := 0; i <= len(s)-span; i++ {
		p := int64(1) // reset current product
		for j := 0; j < span; j++ {
			b := s[i+j]
			if b < '0' || b > '9' {
				return 0, errors.New("invalid symbol in input string")
			}
			p *= int64(b - '0')
		}
		if p > maxP {
			maxP = p
		}
	}
	return
}
