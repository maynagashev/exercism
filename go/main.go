package main

import (
	"fmt"
)

func main() {
	Valid("2323 2005 7766 3554")
}

func Valid(s string) bool {
	// decode digits from string
	var digits []int
	//for _, r := range s {
	//	if unicode.IsDigit(r) {
	//		fmt.Printf("%T %+v\n", r-'0', r-'0')
	//		digits = append(digits, int(r-'0'))
	//	}
	//}

	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]
		switch {
		case c == ' ':
			continue
		case c >= '0' && c <= '9':
			digits = append(digits, int(c-'0'))
		default:
			fmt.Println(digits)
			return false
		}
	}

	//validate
	fmt.Println(digits)
	return true
}
