// Package proverb for generating relevant proverb from a list of inputs
package proverb

import "fmt"

// Proverb generates proverb "For Want of a Nail" from a given list `["nail", "shoe", ..., "kingdom"]
func Proverb(rhyme []string) (result []string) {

	// on empty rhyme just return empty result
	if len(rhyme) < 1 {
		return
	}

	// first processing sentences with multiple words
	for i := 0; i < len(rhyme)-1; i++ {
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1]))
	}

	// processing last sentence if rhyme is not empty
	result = append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))

	return
}
