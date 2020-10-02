// +build bonus

package robotname

import (
	"testing"
	"time"
	"fmt"
)

var maxNames = 26 * 26 * 10 * 10 * 10

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func TestCollisions(t *testing.T) {
	// Test uniqueness for new robots.
	for i := len(seen); i < maxNames; i++ {
		start := time.Now()
		_ = New().getName(t, false)
		if i > maxNames-10 || i<10{
			fmt.Printf("iteration %d took %v\n", i, time.Since(start))
		}
	}

	// Test that names aren't recycled either.
	r := New()
	for i := len(seen); i < maxNames; i++ {
		r.Reset()
		_ = r.getName(t, false)
	}

	// Test that name exhaustion is handled more or less correctly.
	_, err := New().Name()
	if err == nil {
		t.Fatalf("should return error if namespace is exhausted")
	}
}
