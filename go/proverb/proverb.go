// Package proverb for generating relevant proverb from a list of inputs
package proverb

import "fmt"

// Proverb generates proverb "For Want of a Nail" from a slice with rhyme words
func Proverb(rhyme []string) (result []string) {
	wCount := len(rhyme)

	// first processing sentences with multiple words
	if wCount > 1 {
		for i := 0; i < wCount-1; i++ {
			result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
		}
	}

	// processing last sentence if rhyme is not empty
	if wCount > 0 {
		result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
	}

	return
}
