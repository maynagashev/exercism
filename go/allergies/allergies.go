/*
Package allergies contains functions for manage allergy scores.

Benchmarks:

1. Map
BenchmarkAllergies
BenchmarkAllergies-12             458064              2620 ns/op            1008 B/op         17 allocs/op
BenchmarkAllergicTo
BenchmarkAllergicTo-12          13823426                79.8 ns/op             0 B/op          0 allocs/op

2. substance struct
BenchmarkAllergies
BenchmarkAllergies-12            1765278               692 ns/op             112 B/op          5 allocs/op
BenchmarkAllergicTo
BenchmarkAllergicTo-12           8414517               143 ns/op               0 B/op          0 allocs/op

3. substance struct with index map
BenchmarkAllergies
BenchmarkAllergies-12            2083118               563 ns/op             112 B/op          5 allocs/op
BenchmarkAllergicTo
BenchmarkAllergicTo-12          12922242                82.8 ns/op             0 B/op          0 allocs/op

Tags: bitwise, search, maps, structs
*/
package allergies

import "errors"

type substance struct {
	name  string
	score uint
}

var (
	substances = []substance{
		{"eggs", 1},
		{"peanuts", 2},
		{"shellfish", 4},
		{"strawberries", 8},
		{"tomatoes", 16},
		{"chocolate", 32},
		{"pollen", 64},
		{"cats", 128},
	}
	index = makeIndex(substances)
)

// Allergies returns list of substances which was "packed" in given allergy score
func Allergies(score uint) (substanceNames []string) {
	for _, substance := range substances {
		if AllergicTo(score, substance.name) {
			substanceNames = append(substanceNames, substance.name)
		}
	}
	return
}

// AllergicTo determines if given overall score contains specific substance score
func AllergicTo(score uint, substanceName string) bool {
	substanceScore, err := findScore(substanceName)
	return err == nil && score&substanceScore == substanceScore
}

// findScore searches score of substanceName (in source array or in index map, depends on implementation)
func findScore(substanceName string) (uint, error) {
	i, found := index[substanceName] // searching substance in map
	if !found {
		return 0, errors.New(" substance not found")
	}
	return substances[i].score, nil
}

// makeIndex creates search index for fast search substance by name
func makeIndex(substances []substance) map[string]int {
	m := make(map[string]int, len(substances))
	for i, v := range substances {
		m[v.name] = i
	}
	return m
}
