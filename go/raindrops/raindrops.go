/*
Package raindrops for working with raindrop sounds
*/
package raindrops

import "strconv"

// Convert converts integer to raindrops sounds
func Convert(x int) (result string) {
	if x%3 == 0 {
		result += "Pling"
	}
	if x%5 == 0 {
		result += "Plang"
	}
	if x%7 == 0 {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(x)
	}
	return
}
