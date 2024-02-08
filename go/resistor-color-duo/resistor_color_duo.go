package resistorcolorduo

var colors = []string{
	"black",
	"brown",
	"red",
	"orange",
	"yellow",
	"green",
	"blue",
	"violet",
	"grey",
	"white",
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	return ColorCode(colors[0])*10 + ColorCode(colors[1])
}

// ColorCode should return the index of a color.
func ColorCode(color string) int {
	for i, c := range colors {
		if c == color {
			return i
		}
	}
	return -1
}
