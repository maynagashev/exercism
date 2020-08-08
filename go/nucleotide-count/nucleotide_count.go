/*
Package dna contains DNA type and method DNA.Counts for it
*/
package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]uint32

// DNA is a list of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (dna DNA) Counts() (Histogram, error) {
	h := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for i, r := range dna {
		if _, ok := h[r]; !ok {
			return h, fmt.Errorf("invalid nucleotide: %c on position: %d", r, i+1)
		}
		h[r]++
	}

	return h, nil
}
