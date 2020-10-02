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

const max = 26 * 26 * 10 * 10 * 10

var used = make(map[string]bool)
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

// Name returns current Robot name or generates new unique name (not used before).
func (r *Robot) Name() (name string, err error) {
	if r.name != "" {
		return r.name, err
	}
	if len(used) >= max {
		return "", fmt.Errorf("generated %d names, no more variants available", len(used))
	}
	r.name = newName()
	for used[r.name] {
		r.name = newName()
	}
	used[r.name] = true
	return r.name, err
}

// Reset clears current name of Robot
func (r *Robot) Reset() {
	r.name = ""
}

// newName returns new random name for the Robot
func newName() string {
	r1 := random.Intn(26) + 'A'
	r2 := random.Intn(26) + 'A'
	num := random.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
