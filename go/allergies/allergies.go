/*
Package allergies contains functions for manage allergy scores.

Benchmarks:

BenchmarkAllergies
BenchmarkAllergies-12             458064              2620 ns/op            1008 B/op         17 allocs/op
BenchmarkAllergicTo
BenchmarkAllergicTo-12          13823426                79.8 ns/op             0 B/op          0 allocs/op
*/
package allergies

import "sort"

var mapBySubstance = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Allergies returns list of substances which was "packed" in given allergy score
func Allergies(score uint) (substances []string) {
	for _, substance := range sortedSubstances(mapBySubstance) {
		if AllergicTo(score, substance) {
			substances = append(substances, substance)
		}
	}
	return
}

// AllergicTo determines if given overall score contains specific substance score
func AllergicTo(score uint, substance string) bool {
	substanceScore := mapBySubstance[substance]
	return score&substanceScore == substanceScore // true if given score "contains" substance score
}

func sortedSubstances(m map[string]uint) []string {
	scores := make([]int, 0, len(m))
	substances := make([]string, 0, len(m))
	mapByScore := make(map[uint]string, len(m))

	for name, score := range m {
		mapByScore[score] = name
		scores = append(scores, int(score))
	}

	sort.Ints(scores)

	for _, score := range scores {
		substances = append(substances, mapByScore[uint(score)])
	}
	return substances
}
