/*
Package phonenumber contains functions for handle phone numbers.

Benchmarks:
1. onlyDigitsRegexp
BenchmarkNumber
BenchmarkNumber-12                 78864             13950 ns/op            1466 B/op         97 allocs/op
BenchmarkAreaCode
BenchmarkAreaCode-12               87020             13914 ns/op            1468 B/op         97 allocs/op
BenchmarkFormat
BenchmarkFormat-12                 63368             19038 ns/op            2639 B/op        169 allocs/op

2. onlyDigitsMap
BenchmarkNumber
BenchmarkNumber-12                635520              1935 ns/op             608 B/op         27 allocs/op
BenchmarkAreaCode
BenchmarkAreaCode-12              608353              1889 ns/op             608 B/op         27 allocs/op
BenchmarkFormat
BenchmarkFormat-12                200294              5922 ns/op            1760 B/op         99 allocs/op

3. onlyDigitsFor & builder
BenchmarkNumber
BenchmarkNumber-12                958804              1248 ns/op             528 B/op         31 allocs/op
BenchmarkAreaCode
BenchmarkAreaCode-12              905804              1262 ns/op             528 B/op         31 allocs/op
BenchmarkFormat
BenchmarkFormat-12                228822              5212 ns/op            1680 B/op        103 allocs/op
*/
package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Number extracts number from input
func Number(input string) (res string, err error) {
	// clean up
	digits := onlyDigitsFor(input)

	// special case
	if len(digits) == 11 && digits[0] == '1' {
		digits = digits[1:]
	}

	// validation
	switch {
	case len(digits) != 10:
		return digits, errors.New("invalid length")
	case digits[0] == '1', digits[0] == '0':
		return digits, errors.New("invalid area code")
	case digits[3] == '1', digits[3] == '0':
		return digits, errors.New("invalid exchange code")
	}

	return digits, nil
}

// AreaCode extracts area code from input
func AreaCode(number string) (res string, err error) {
	num, err := Number(number)
	return num[:3], err
}

// Format formats phone number
func Format(number string) (res string, err error) {
	num, err := Number(number)
	return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), err
}

func onlyDigitsFor(input string) string {
	var b strings.Builder
	b.Grow(len(input))
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c >= '0' && c <= '9' {
			b.WriteByte(c)
		}
	}
	return b.String()
}

func onlyDigitsMap(input string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r >= '0' && r <= '9':
			return r
		default:
			return -1
		}
	}, input)
}

var notDigits = regexp.MustCompile(`[^\d]`)
func onlyDigitsRegexp(input string) string {
	return string(notDigits.ReplaceAll([]byte(input), []byte("")))
}