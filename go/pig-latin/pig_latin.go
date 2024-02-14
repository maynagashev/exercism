package piglatin

import (
	"strings"
)

func Sentence(sentence string) string {
	words := strings.Fields(sentence)
	pigLatinWords := make([]string, len(words))

	for i, word := range words {
		pigLatinWords[i] = translateWord(word)
	}

	return strings.Join(pigLatinWords, " ")
}

func translateWord(word string) string {
	firstLetter := string(word[0])
	firstTwoLetters := word[0:2]

	// Rule 1. Words beginning with a vowel or xr or yt
	if isVowel(firstLetter) || (len(word) > 2 && firstTwoLetters == "xr" || firstTwoLetters == "yt") {
		return word + "ay"
	}

	// Rule 3. Consonant followed by qu
	if len(word) > 2 && isConsonant(firstLetter) && word[1:3] == "qu" {
		return word[3:] + firstLetter + "quay"
	}

	// Rule 4. Consonant sound or cluster with following y
	if len(word) > 2 && isConsonant(firstLetter) && isConsonant(word[1:2]) && word[2:3] == "y" {
		return word[2:] + firstTwoLetters + "ay"
	}
	if len(word) > 2 && isConsonant(firstLetter) && word[1:2] == "y" {
		return word[1:] + firstLetter + "ay"
	}

	// Rule 2. Consonant sound and special consonant cases
	if len(word) > 2 && (word[0:3] == "thr" || word[0:3] == "sch") {
		return word[3:] + word[0:3] + "ay"
	}
	if firstTwoLetters == "ch" || firstTwoLetters == "st" || firstTwoLetters == "qu" || firstTwoLetters == "th" {
		return word[2:] + firstTwoLetters + "ay"
	}
	if isConsonant(firstLetter) {
		return word[1:] + firstLetter + "ay"
	}

	return word + "ay"
}

func isConsonant(letter string) bool {
	return strings.Contains("bcdfghjklmnpqrstvwxyz", letter)
}

func isVowel(letter string) bool {
	return strings.Contains("aeiou", letter)
}
