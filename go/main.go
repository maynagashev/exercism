package main

import (
	"fmt"
	"sort"
)

var m = map[int]string{
	1:   "eggs",
	2:   "peanuts",
	4:   "shellfish",
	8:   "strawberries",
	16:  "tomatoes",
	32:  "chocolate",
	64:  "pollen",
	128: "cats",
}

var m2 = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}
// list
var listTests = []struct {
	description string
	score       uint
	expected    []string
}{
	{"no allergies at all", 0, []string{}},
	{"allergic to just eggs", 1, []string{"eggs"}},
	{"allergic to just peanuts", 2, []string{"peanuts"}},
	{"allergic to just strawberries", 8, []string{"strawberries"}},
	{"allergic to eggs and peanuts", 3, []string{"eggs", "peanuts"}},
	{"allergic to more than eggs but not peanuts", 5, []string{"eggs", "shellfish"}},
	{"allergic to lots of stuff", 248, []string{"strawberries", "tomatoes", "chocolate", "pollen", "cats"}},
	{"allergic to everything", 255, []string{"eggs", "peanuts", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}},
	{"ignore non allergen score parts", 509, []string{"eggs", "shellfish", "strawberries", "tomatoes", "chocolate", "pollen", "cats"}},
}


func main() {

	fmt.Printf("%+q", sortedSubstances(m2))

}

func sortedSubstances(m map[string]uint) []string {
	mapByScore := make(map[uint]string)
	scores := make([]int, 0, len(m))
	substances := make([]string, 0, len(m))

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