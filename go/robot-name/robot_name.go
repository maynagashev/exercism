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
var m = make(map[int]bool)

// Name returns current Robot name
// If name is empty, then will be generated new name
func (r *Robot) Name() (string, error) {
	if r.name == "" {
		var err error
		r.name, err = newName()
		return r.name, err
	}
	return r.name, nil
}

// Reset clears current name of Robot
func (r Robot) Reset() {
	r.name = ""
}

// newName generates new random name for Robot
func newName() (string, error) {
	for i := 0; i < maxNameVariants; i++ {
		num := rand.Intn(maxNameVariants)
		runes := make([]rune, 5)
		runes[0] = rune('A' + num/1000/26%26)
		runes[1] = rune('A' + num/1000%26)
		runes[2] = rune('0' + num/100%10)
		runes[3] = rune('0' + num/10%10)
		runes[4] = rune('0' + num%10)
		name := string(runes)
		fmt.Println(num, "=", name, m)
		if _, ok:= m[num]; !ok {
			m[num] = true
			return name, nil
		}
	}
	return "", errors.New("all names occupied")
}
