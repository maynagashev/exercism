/*
Package cipher - implements a simple shift cipher like Caesar and a more secure
substitution cipher.
*/
package cipher

import (
	"strings"
	"unicode"
)

// Cipher interface for different implementations
type Cipher interface {
	Encode(string) string
	Decode(string) string
}

// cipher struct with encoding maps
type cipher struct {
	encodeMap map[rune]rune
	decodeMap map[rune]rune
}

// convert converts given string using given conversion map
// runes not presented in map - will be just excluded from the results
func (c cipher) convert(s string, conversionMap map[rune]rune) string {
	b := strings.Builder{}
	for _, r := range s {
		if converted, ok := conversionMap[unicode.ToLower(r)]; ok {
			b.WriteRune(converted)
		}
	}
	return b.String()
}

// Encode converts given string using cipher.encodeMap.
func (c cipher) Encode(s string) string {
	return c.convert(s, c.encodeMap)
}

// Encode converts given string using cipher.decodeMap.
func (c cipher) Decode(s string) string {
	return c.convert(s, c.decodeMap)
}

// NewCaesar generates conversion maps for simple caesar cipher with shift=3.
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift generates conversion maps for simple caesar cipher with specified shift.
func NewShift(shift int) Cipher {
	var c cipher

	abcLen := 'z' - 'a'
	if shift == 0 || int32(shift) < (-abcLen) || int32(shift) > abcLen {
		return nil
	}

	c.encodeMap = make(map[rune]rune, abcLen)
	c.decodeMap = make(map[rune]rune, abcLen)

	for i := 'a'; i <= 'z'; i++ {
		j := i + int32(shift)
		// handle roll over
		if j > 'z' {
			j -= abcLen + 1
		}
		if j < 'a' {
			j += abcLen + 1
		}
		c.encodeMap[i] = j
		c.decodeMap[j] = i
	}

	return c
}

func NewVigenere(k string) Cipher {
	var c cipher
	return c
}
