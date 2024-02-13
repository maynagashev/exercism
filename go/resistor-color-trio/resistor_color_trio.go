package resistorcolortrio

import (
	"fmt"
	"math"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	colorValues := map[string]int{"black": 0, "brown": 1, "red": 2, "orange": 3, "yellow": 4, "green": 5, "blue": 6, "violet": 7, "grey": 8, "white": 9}
	value := colorValues[colors[0]]*10 + colorValues[colors[1]]
	zeros := int(math.Pow(10, float64(colorValues[colors[2]])))
	ohms := value * zeros
	switch {
	case ohms >= 1000000000:
		return fmt.Sprintf("%d gigaohms", ohms/1000000000)
	case ohms >= 1000000:
		return fmt.Sprintf("%d megaohms", ohms/1000000)
	case ohms >= 1000:
		return fmt.Sprintf("%d kiloohms", ohms/1000)
	default:
		return fmt.Sprintf("%d ohms", value*zeros)
	}
}
