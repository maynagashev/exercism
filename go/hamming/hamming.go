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

import "errors"

// Distance counting different symbols on the same positions in two equal length strings
func Distance(a, b string) (int, error) {

	runesA, runesB, distance := []rune(a), []rune(b), 0

	if len(runesA) != len(runesB) {
		return 0, errors.New("sequences is not equal in length")
	}

	for pos := range runesA {
		if runesA[pos] != runesB[pos] {
			distance++
		}
	}

	return distance, nil
}
