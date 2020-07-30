/*
Raindrops

Your task is to convert a number into a string that contains raindrop sounds corresponding to certain potential factors.
A factor is a number that evenly divides into another number, leaving no remainder. The simplest way to test if a one
number is a factor of another is to use the [modulo operation](https://en.wikipedia.org/wiki/Modulo_operation).

The rules of `raindrops` are that if a given number:

- has 3 as a factor, add 'Pling' to the result.
- has 5 as a factor, add 'Plang' to the result.
- has 7 as a factor, add 'Plong' to the result.
- _does not_ have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
*/
package raindrops

import "strconv"

// isFactor determines whether specified number is a factor of another
func isFactor(x int, y int) bool {
	return x%y == 0
}

// Convert converts integer to raindrops sounds
func Convert(x int) string {
	var result string
	if isFactor(x, 3) {
		result += "Pling"
	}
	if isFactor(x, 5) {
		result += "Plang"
	}
	if isFactor(x, 7) {
		result += "Plong"
	}

	if result == "" {
		result = strconv.Itoa(x)
	}
	return result
}
