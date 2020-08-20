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


*/
package main

import (
	"fmt"
	"errors"
)

func main() {
	LargestSeriesProduct("(0123456789", 2)
}

//LargestSeriesProduct calculates the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(s string, span int) (product int64, err error) {
	if span > len(s) || span < 0 {
		return 0, errors.New("span is out of bounds")
	}
	for i, p := 0, int64(1); i < len(s)-span+1; i++ {
		fmt.Println(p)
		for j := 0; j < span; j++ {
			fmt.Println(i, j, s[i+j])
			if s[i+j] < '0' || s[i+j] > '9' {
				return 0, errors.New("invalid symbol in string")
			}
			p *= int64(s[i+j] - '0')
		}
		if p > product {
			product = p
		}
	}
	return
}

