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
func newName() (name string) {
	rnd := rand.Intn(maxNameVariants)
	rndDigit := rand.Intn(1000)
	fmt.Println("newName() maxNameVariants: ", maxNameVariants, "rnd (full): ", rnd, (rnd+26)%rnd, "rndDigit:", rndDigit)
	d1 := rndDigit/100 - rndDigit/100/10*10
	d2 := rndDigit/10 - rndDigit/10/10*10
	d3 := rndDigit - rndDigit/10*10
	name = "AA" + string(rune('0'+d1)) + string(rune('0'+d2)) + string(rune('0'+d3))
	fmt.Println(name)
	return name
}
