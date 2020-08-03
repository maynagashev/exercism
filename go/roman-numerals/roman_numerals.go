/*
Package romannumerals contains functions for work with Roman Numerals
*/
package romannumerals

import (
	"fmt"
)

var (
	units     = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands = []string{"", "M", "MM", "MMM"}
)

// ToRomanNumeral converts arabic number to roman numeral
func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3000 {
		return "", fmt.Errorf("out of range: %v", n)
	}
	return thousands[n/1000] + hundreds[n%1000/100] + tens[n%100/10] + units[n%10], nil
}
