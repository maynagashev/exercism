/*
Package anagram

An anagram is a rearrangement of letters to form a new word.
Given a word and a list of candidates, select the sublist of anagrams of the given word.
*/
package main

import (
	"fmt"
	"strings"
)

func main() {

	Detect("allergy", []string{
		"gallery",
		"ballerina",
		"regally",
		"clergy",
		"largely",
		"leading"})
}
// Detect selects the sublist of anagrams of the given word
func Detect(word string, candidates []string) (sublist []string) {
	for _, cWord := range candidates {
		if isAnagram(word, cWord) {
			sublist = append(sublist, cWord)
		}
	}

	return
}

func isAnagram(word string, word2 string) bool {
	word, word2 = strings.ToLower(word), strings.ToLower(word2)

	// check: !same word && have same len
	if word == word2 || len(word) != len(word2) {
		return false
	}

	usedI2 := map[int]bool{}
	fmt.Println(word, word2)
	// searching runes of the first word in a second word
	for _, r1 := range word {
		fmt.Println(r1, string(r1))
		r1Found := false
		// scanning runes of second word
		for i2, r2 := range word2 {
			fmt.Println("\t r2:", r2, string(r2), usedI2)
			_, isUsed := usedI2[i2]
			if r1 == r2 && !isUsed {
				r1Found = true
				usedI2[i2] = true
				break;
			}
		}
		if !r1Found {
			return false
		}
	}

	return true
}
