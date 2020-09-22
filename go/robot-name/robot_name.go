package robotname

import "fmt"

type Robot struct {
}

const max = 26 * 26 * 10 * 10 * 10

func (*Robot) Name() (string, error) {
	fmt.Println("max: ", max)
	return "AA222", nil
}

func (r Robot) Reset() {

}
