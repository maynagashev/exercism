package elon

import "fmt"

// Drive define the 'Drive()' method
func (c *Car) Drive() {
	var newBattery = c.battery - c.batteryDrain
	if newBattery < 0 {
		return
	}
	c.battery = newBattery
	c.distance += c.speed
}

// DisplayDistance define the 'DisplayDistance() string' method
func (c *Car) DisplayDistance() string {
	return "Driven " + fmt.Sprint(c.distance) + " meters"
}

// DisplayBattery define the 'DisplayBattery() string' method
func (c *Car) DisplayBattery() string {
	return "Battery at " + fmt.Sprint(c.battery) + "%"
}

// CanFinish define the 'CanFinish(int) bool' method
func (c *Car) CanFinish(trackDistance int) bool {
	return trackDistance <= c.battery/c.batteryDrain*c.speed
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
