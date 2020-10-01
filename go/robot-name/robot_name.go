package robotname

import (
	"fmt"
	"math/rand"
	"time"
)

// Robot with name
type Robot struct {
	name string
}

const maxNameVariants = 26 * 26 * 10 * 10 * 10

var variants []int
var nextVariant int

func init() {
	Permutate()
}

// Permutate generates a pseudo-random permutation of the integers [0,maxNameVariants)
func Permutate() {
	rand.Seed(time.Now().UnixNano())
	variants = rand.Perm(maxNameVariants)
}

// Name returns current Robot name
// If name is empty, then will be generated new name
func (r *Robot) Name() (name string, err error) {
	if r.name == "" {
		r.name, err = newName()
	}
	return r.name, err
}

// Reset clears current name of Robot
func (r *Robot) Reset() {
	r.name = ""
}

// newName returns new random name for the Robot which not used before (with guarantee).
// Uses prefilled variants slice of randomly permutated integers.
func newName() (string, error) {
	if nextVariant >= len(variants) {
		return "", fmt.Errorf("generated %d names, no more variants available", nextVariant)
	}
	num := variants[nextVariant]
	nextVariant++

	runes := make([]rune, 5)
	runes[0] = rune('A' + num/1000/26%26)
	runes[1] = rune('A' + num/1000%26)
	runes[2] = rune('0' + num/100%10)
	runes[3] = rune('0' + num/10%10)
	runes[4] = rune('0' + num%10)
	return string(runes), nil
}
