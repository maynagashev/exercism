/*
Package scale presenting musical scale generator

A, A#, B, C, C#, D, D#, E, F, F#, G, G#
A, Bb, B, C, Db, D, Eb, E, F, Gb, G, Ab

The collection of notes in these scales is written with either sharps or flats, depending on the tonic. Here is a list of which are which:
 - No Sharps or Flats: C major a minor
 - Use Sharps: G, D, A, E, B, F# major e, b, f#, c#, g#, d# minor
 - Use Flats: F, Bb, Eb, Ab, Db, Gb major d, g, c, f, bb, eb minor

The diatonic scales, and all other scales that derive from the
chromatic scale, are built upon intervals. An interval is the space
between two pitches.

The simplest interval is between two adjacent notes, and is called a
"half step", or "minor second" (sometimes written as a lower-case "m").
The interval between two notes that have an interceding note is called
a "whole step" or "major second" (written as an upper-case "M"). The
diatonic scales are built using only these two intervals between
adjacent notes.
*/
package scale

import "strings"

// Scale generates musical scale from given tonic and with given interval
func Scale(tonic string, interval string) (scale []string) {
	scale = []string{}
	switch tonic {
	case "C", "G", "D", "A", "E", "B", "F#", "a", "e", "b", "f#", "c#", "g#", "d#":
		scale = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		scale = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	}

	tonic = strings.Title(tonic)
	for i, elem := range scale {
		if elem == tonic {
			scale = append(scale[i:], scale[:i]...)
			break
		}
	}

	if interval == "" {
		return scale
	}

	var partialScale []string
	stepSize := map[string]int{"m": 1, "M": 2, "A": 3}
	i := 0
	for _, diff := range strings.Split(interval, "") {
		if step, ok := stepSize[diff]; ok {
			partialScale = append(partialScale, scale[i%len(scale)])
			i += step
		}
	}

	return partialScale
}
