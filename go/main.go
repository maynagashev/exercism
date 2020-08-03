package main

import (
	"fmt"
)

func main() {
	fmt.Println(2136)
	fmt.Println(ToRomanNumeral(2136))
}

var (
	units     = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands = []string{"", "M", "MM", "MMM"}
)

func ToRomanNumeral(n int) (string, error) {
	if n < 1 || n > 3000 {
		return "", fmt.Errorf("out of range: %v", n)
	}
	fmt.Println(float64(n)/1000, n%1000/100, n%100/10, n%10)
	return thousands[n/1000] + "." + hundreds[n%1000/100] + "." + tens[n%100/10] + "." + units[n%10], nil
}
