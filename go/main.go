package main

import (
	"fmt"
	"unicode"
)

func main() {
	IsPangram("Five quacking Zephyrs jolt my wax bed.")
}

// IsPangram determines if the given string contains every letter
func IsPangram(s string) bool {
	var bs int32
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v %c %v %v %c %+v %T\n", unicode.ToUpper(rune(i)), s[i], s[i], 0xdf, s[i] & 0xdf, s[i] & 0xdf, (s[i] & 0xdf))
		c := (s[i] & 0xdf) - 'A'
		if c > 25 || c < 0 {
			continue
		}
		bs |= 1 << c
	}
	return bs == 0x3ffffff
}
