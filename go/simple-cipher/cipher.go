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

// caesarCipher struct with encoding maps
type caesarCipher struct {
	encodeMap map[rune]rune
	decodeMap map[rune]rune
}

// Encode converts given string using caesarCipher.encodeMap.
func (c caesarCipher) Encode(s string) string {
	return c.convert(s, c.encodeMap)
}

// Encode converts given string using caesarCipher.decodeMap.
func (c caesarCipher) Decode(s string) string {
	return c.convert(s, c.decodeMap)
}

// convert converts given string using given conversion map
// runes not presented in map - will be just excluded from the results
func (c caesarCipher) convert(s string, conversionMap map[rune]rune) string {
	b := strings.Builder{}
	for _, r := range s {
		if converted, ok := conversionMap[unicode.ToLower(r)]; ok {
			b.WriteRune(converted)
		}
	}
	return b.String()
}

// NewCaesar generates conversion maps for simple caesar cipher with shift=3.
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift generates conversion maps for simple caesar cipher with specified shift.
func NewShift(shift int) Cipher {
	var c caesarCipher

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

type vigenereCipher struct {
	keyword string
}

// Encode encodes string using Vigenere cipher with encapsulated keyword
func (c vigenereCipher) Encode(s string) string {
	// todo
	return s
}

// Encode decodes string using Vigenere cipher with encapsulated keyword
func (c vigenereCipher) Decode(s string) string {
	// todo
	return s
}

// NewVigenere generates conversion maps for encoding text by using a series of interwoven Caesar ciphers,
// based on the letters of a given keyword.
func NewVigenere(keyword string) Cipher {
	return vigenereCipher{keyword}
}
