/*
Package prime for working with prime numbers.

Benchmarks:
BenchmarkNth
BenchmarkNth-12              103          11492036 ns/op               0 B/op          0 allocs/op

Tags: primes.
*/
package prime

import "math"

// isPrime determines if given value is prime without using sieves
func isPrime(v int) bool {
	// filter all even variants
	if v%2 == 0 {
		return v == 2
	}

	// search any odd divisors lower than root
	root := int(math.Sqrt(float64(v)))
	for t := 3; t <= root; t += 2 {
		if v%t == 0 {
			return false
		}
	}
	return true
}

// Nth resolves n-th prime
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	// iterating continuously until N primes are resolved, starting from the first prime (2)
	for i, p := 0, 2; ; p++ {
		if isPrime(p) {
			if i++; i == n { // stop when found N-th prime
				return p, true
			}
		}
	}
}
