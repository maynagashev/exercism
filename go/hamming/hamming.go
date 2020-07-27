/*
Package hamming contains method for calculating Hamming Distance between two DNA strands (presented as strings).

We read DNA using the letters C,A,G and T. Two strands might look like this:

    GAGCCTACTAACGGGAT
    CATCGTAATGACGGCCT
    ^ ^ ^  ^ ^    ^^

They have 7 differences, and therefore the Hamming Distance is 7.

The Hamming Distance is useful for lots of things in science, not just biology, so it's a nice phrase to be familiar with :)
*/
package hamming

import "github.com/pkg/errors"

// Distance counting different symbols on the same positions in two equal length strings
func Distance(a, b string) (int, error) {

	distance := 0

	if len(a) != len(b) {
		return distance, errors.New("sequences is not equal in length")
	}

	for pos := range a {
		if a[pos] != b[pos] {
			distance++
		}
	}

	return distance, nil
}
