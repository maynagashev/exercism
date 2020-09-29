package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot with name
type Robot struct {
	name string
}

const maxNameVariants = 26 * 26 * 10 * 10 * 10

var generatedNames = make(map[string]bool)

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

// newName generates new random name for Robot
func newName() (string, error) {
	for i := 0; i < maxNameVariants; i++ {
		name := randomName2()
		if _, ok := generatedNames[name]; !ok {
			generatedNames[name] = true
			return name, nil
		}
	}
	return "", errors.New("all names occupied")
}

func randomName() string {
	num := rand.Intn(maxNameVariants)
	runes := make([]rune, 5)
	runes[0] = rune('A' + num/1000/26%26)
	runes[1] = rune('A' + num/1000%26)
	runes[2] = rune('0' + num/100%10)
	runes[3] = rune('0' + num/10%10)
	runes[4] = rune('0' + num%10)
	return string(runes)
}

func randomName2() string {
	r1 := rand.Intn(26) + 'A'
	r2 := rand.Intn(26) + 'A'
	num := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
