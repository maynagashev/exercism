/*
Package luhn contains the Luhn validation function for number sequences

1. The first step of the Luhn algorithm is to double every second digit,
starting from the right.
2. If doubling the number results in a number greater than 9 then subtract 9
from the product.
3. Sum all of the digits.
4. If the sum is evenly divisible by 10, then the number is valid.

*/
package luhn

// Valid determines whether or not given number is valid per the Luhn formula.
// Allowed only digits and spaces.
// Strings of length 1 or less are not valid.
// Example: 4539 1488 0343 6467
func Valid(s string) bool {
	var sum, i, d, di int

	// Extracting digits from end without decoding Unicode runes
	for i = len(s) - 1; i >= 0; i-- {
		c := s[i]
		switch {
		// Skipping spaces
		case c == ' ':
			continue
		// Processing only digits
		case c >= '0' && c <= '9':
			d = int(c - '0')
			// Doubling every second digit
			if di%2 == 1 {
				d *= 2
			}
			// Subtract 9
			if d > 9 {
				d -= 9
			}
			sum += d
			di++
		// If faced any invalid characters in input string, then invalidate result completely
		default:
			return false
		}
	}

	return di > 1 && sum%10 == 0
}
