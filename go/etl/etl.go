/*
Package etl contains transforming method for some Scrabble scores from a legacy system
*/
package etl

import "strings"

// Transform transforms some Scrabble scores from a legacy system
func Transform(input map[int][]string) map[string]int {
	output := make(map[string]int)
	for score, letters := range input {
		for _, s := range letters {
			output[strings.ToLower(s)] = score
		}
	}

	return output
}
