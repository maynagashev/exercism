package luhn

import (
	"fmt"
	"unicode"
)

func Valid(s string) bool {
	// decode digits from string
	var digits []int
	for _, r := range s {
		if unicode.IsDigit(r) {
			digits = append(digits, int(r))
		}
	}

	//validate
	fmt.Println(digits)
}