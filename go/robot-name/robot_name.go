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

var generatedNames = make(map[string]bool)

func init() {
	rand.Seed(time.Now().UnixNano())
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

// newName returns new random name for the Robot which not used before.
func newName() (string, error) {
	for len(generatedNames) < maxNameVariants {
		r1 := rand.Intn(26) + 'A'
		r2 := rand.Intn(26) + 'A'
		num := rand.Intn(1000)
		name := fmt.Sprintf("%c%c%03d", r1, r2, num)
		if _, seen := generatedNames[name]; !seen {
			generatedNames[name] = true
			return name, nil
		}
	}
	return "", fmt.Errorf("generated %d names, no more variants available", len(generatedNames))
}
