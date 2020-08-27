/*
Package sieve uses the Sieve of Eratosthenes algorithm to find all the primes from 2 up to a given number.

Algorithm:

- take the next available unmarked number in your list (it is prime)
- mark all the multiples of that number (they are not prime - composite)

Benchmarks:

1. primes := make([]int, 0)
BenchmarkSieve
BenchmarkSieve-12    	  269534	      3897 ns/op	    5336 B/op	      22 allocs/op

2. primes := make([]int, 0, limit)
BenchmarkSieve
BenchmarkSieve-12    	  382113	      3019 ns/op	    9472 B/op	      10 allocs/op

3. primes := make([]int, 0, limit*4/10)
BenchmarkSieve
BenchmarkSieve-12         522270              2299 ns/op            4432 B/op         10 allocs/op

4. iterate only through odd numbers
BenchmarkSieve
BenchmarkSieve-12         696003              1707 ns/op            4432 B/op         10 allocs/op

*/
package sieve

// Sieve finds all the primes from 2 up to a given number.
func Sieve(limit int) []int {
	composite := make([]bool, limit+1)
	primes := make([]int, 0, limit*4/10)

	if limit > 1 {
		primes = append(primes, 2)
	}
	// further check only "odd" numbers
	for i := 3; i <= limit; i += 2 {
		if composite[i] {
			continue
		}
		primes = append(primes, i)
		for j := i * 2; j <= limit; j += i {
			composite[j] = true
		}
	}
	return primes
}
