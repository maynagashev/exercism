/*
Package anagram contains detection methods for anagrams

An anagram is a rearrangement of letters to form a new word.
Given a word and a list of candidates, select the sublist of anagrams of the given word.
*/
package anagram

import (
	"strings"
)

// Detect selects the sublist of anagrams of the given word
func Detect(word string, candidates []string) (sublist []string) {
	for _, candidate := range candidates {
		if IsAnagram(word, candidate) {
			sublist = append(sublist, candidate)
		}
	}

	return
}

// IsAnagram determines if a candidate word is an anagram of given word
func IsAnagram(word string, candidate string) bool {

	// all matching case insensitive
	word, candidate = strings.ToLower(word), strings.ToLower(candidate)

	// check: !same words && has same len
	if word == candidate || len(word) != len(candidate) {
		return false
	}

	// memory for used chars in candidate
	mem := map[int]bool{}

	// searching runes of the word in the candidate
	for _, r1 := range word {
		isFound := false
		// scanning runes of the candidate
		for i2, r2 := range candidate {
			_, isUsed := mem[i2]
			if r1 == r2 && !isUsed {
				isFound = true
				mem[i2] = true
				break // each rune searching once in the candidate
			}
		}
		if !isFound {
			return false
		}
	}

	return true
}
