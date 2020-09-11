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

// vigenereCipher struct with specified keyword
type vigenereCipher struct {
	keyword string
}

// NewCaesar generates conversion maps for simple caesar cipher with shift=3.
func NewCaesar() Cipher {
	return NewShift(3)
}

// NewShift generates conversion maps for simple caesar cipher with specified shift.
func NewShift(shift int) Cipher {
	if shift < 0 {
		shift += 26
	}
	return NewVigenere(string('a' + shift)) // new vigenere with repeating constant shift symbol
}

// NewVigenere generates conversion maps for encoding text by using a series of interwoven Caesar ciphers,
// based on the letters of a given keyword.
func NewVigenere(keyword string) Cipher {
	hasValidSymbols := false
	for _, r := range keyword {
		switch {
		case r > 'a' && r <= 'z':
			hasValidSymbols = true
		case r < 'a' || r > 'z':
			return nil // has invalid symbols
		}
	}
	if !hasValidSymbols {
		return nil
	}

	return vigenereCipher{keyword}
}

// Encode encodes string using Vigenere cipher with encapsulated keyword
func (cipher vigenereCipher) Encode(s string) string {
	keywordIndex := 0
	res := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if !unicode.IsLetter(rune(s[i])) {
			continue
		}
		if keywordIndex >= len(cipher.keyword) {
			keywordIndex = 0
		}
		shift := cipher.keyword[keywordIndex] - 'a'
		b := byte(unicode.ToLower(rune(s[i]))) + shift
		if b > 'z' {
			b -= 26
		}
		res.WriteByte(b)
		keywordIndex++
	}
	return res.String()
}

// Encode decodes string using Vigenere cipher with encapsulated keyword
func (cipher vigenereCipher) Decode(s string) string {
	keywordIndex := 0
	res := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if keywordIndex >= len(cipher.keyword) {
			keywordIndex = 0
		}
		shift := cipher.keyword[keywordIndex] - 'a'
		b := byte(unicode.ToLower(rune(s[i]))) - shift
		if b < 'a' {
			b += 26
		}
		res.WriteByte(b)
		keywordIndex++
	}
	return res.String()
}
