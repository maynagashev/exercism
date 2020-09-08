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

// ConcurrentFrequency counts the frequency of each rune in a given texts, running "scanning" freq func concurrently
func ConcurrentFrequency(texts []string) FreqMap {

	// declarations
	res := FreqMap{}
	c := make(chan FreqMap, len(texts))
	freq := func(s string, c chan FreqMap) {
		m := FreqMap{}
		for _, r := range s {
			m[r]++
		}
		c <- m
	}

	// executing goroutines
	for _, t := range texts {
		go freq(t, c)
	}

	// collecting results
	for i := 0; i < len(texts); i++ {
		m := <-c
		for r, n := range m {
			res[r] += n
		}
	}

	return res
}
