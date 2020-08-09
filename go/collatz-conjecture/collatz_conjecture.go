/*
Package collatzconjecture contains function for testing Collatz conjecture

Consider the following operation on an arbitrary positive integer:

If the number is even, divide it by two.
If the number is odd, triple it and add one.

The Collatz conjecture is: This process will eventually reach the number 1, regardless of which positive integer is chosen initially.
*/
package collatzconjecture

import (
	"fmt"
)

// CollatzConjecture calculates the number of steps required to reach 1
func CollatzConjecture(n int) (steps int, err error) {
	for ; n > 1; n = step(n) {
		steps++
	}
	if n != 1 {
		return steps, fmt.Errorf("invalid n=%v on step=%d", n, steps)
	}
	return steps, nil
}

func step(n int) int {
	if n%2 == 0 {
		return n >> 1 // a bit faster than n/2
	}
	return 3*n + 1
}
