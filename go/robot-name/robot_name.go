package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
}

const max = 26 * 26 * 10 * 10 * 10
const period = 26+26+10+10+10

func (*Robot) Name() (string, error) {
	rnd := rand.Intn(max)
	rndDigit := rand.Intn(1000)
	fmt.Println("max: ", max, "rnd: ", rnd, (rnd+26)%rnd, "rndDigit:", rndDigit)
	d1:= rndDigit/100 - rndDigit/100/10*10
	d2 := rndDigit/10 - rndDigit/10/10*10
	d3:= rndDigit - rndDigit/10*10
	res := "AA"+string(rune('0'+d1))+string(rune('0'+d2))+string(rune('0'+d3))
	fmt.Println(res)
	return res, nil
}

func (r Robot) Reset() {

}
