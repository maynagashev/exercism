package robotname

import (
	"fmt"
	"math/rand"
)

// Robot with name
type Robot struct {
	name string
}

const maxNameVariants = 26 * 26 * 10 * 10 * 10

// Name returns current Robot name
// If name is empty, then will be generated new name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = newName()
	}
	return r.name, nil
}

// Reset clears current name of Robot
func (r Robot) Reset() {
	r.name = ""
}

// newName generates new random name for Robot
func newName() string {
	rnd := rand.Intn(maxNameVariants)
	fmt.Println("newName()", maxNameVariants, "rnd (full): ", rnd)
	runes := make([]rune, 5)
	runes[0] = rune('A' + rnd/1000/26%26)
	runes[1] = rune('A' + rnd/1000%26)
	runes[2] = rune('0' + rnd/100%10)
	runes[3] = rune('0' + rnd/10%10)
	runes[4] = rune('0' + rnd%10)
	name := string(runes)
	fmt.Println(runes, name)
	return name
}
