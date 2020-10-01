package robotname

import (
	"fmt"
	"math/rand"
)

// randomNameDisperse generates random name based on single random integer from in [0,maxNameVariants)
func randomNameDisperse() string {
	num := rand.Intn(maxNameVariants)
	runes := make([]rune, 5)
	runes[0] = rune('A' + num/1000/26%26)
	runes[1] = rune('A' + num/1000%26)
	runes[2] = rune('0' + num/100%10)
	runes[3] = rune('0' + num/10%10)
	runes[4] = rune('0' + num%10)
	return string(runes)
}

// randomNameSimple generates random name based on random values for each letter and single random value for numerical part
func randomNameSimple() string {
	r1 := rand.Intn(26) + 'A'
	r2 := rand.Intn(26) + 'A'
	num := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
