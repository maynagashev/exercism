/*
Package letter contains solutions for Parallel Letter Frequency assignment for exercism.io

In this assignment exploring concurrency patterns.
*/
package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of each rune in a given texts, running Frequency concurrently
func ConcurrentFrequency(texts []string) FreqMap {

	c := make(chan FreqMap, len(texts))

	// executing goroutines
	for _, t := range texts {
		go func(s string) {
			c <- Frequency(s)
		}(t)
	}

	// using first result as is (boost: speed 12%, allocations ~24%)
	res := <-c

	// collecting other results
	for i := 1; i < len(texts); i++ {
		for r, n := range <-c {
			res[r] += n
		}
	}
	return res
}
