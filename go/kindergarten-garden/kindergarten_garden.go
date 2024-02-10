package kindergarten

import (
	"fmt"
	"slices"
	"strings"
)

// Define the Garden type here.
type Garden struct {
	row1, row2 []rune // 24 is the maximum number of plants in a row
	children   []string
}

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

// NewGarden returns a new Garden with the given diagram and children.
func NewGarden(diagram string, children []string) (*Garden, error) {
	// check for invalid characters in the diagram
	for _, r := range diagram {
		switch r {
		case 'V', 'C', 'R', 'G', '\n':
			continue
		default:
			return nil, fmt.Errorf("invalid character in diagram: %c", r)
		}
	}

	var g = Garden{
		row1:     make([]rune, len(children)*2),
		row2:     make([]rune, len(children)*2),
		children: make([]string, len(children)),
	}
	// copy the children slice to avoid modifying the caller's slice
	copy(g.children, children)
	slices.Sort(g.children)

	// create a map of the children and check for duplicates
	childMap := make(map[string]bool)
	for _, c := range g.children {
		if childMap[c] {
			return nil, fmt.Errorf("duplicate child: %s", c)
		}
		childMap[c] = true
	}

	// explode the diagram into the two rows
	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 {
		return nil, fmt.Errorf("diagram must have 3 rows")
	}
	if len(rows[1]) != len(rows[2]) || len(rows[1]) != len(children)*2 {
		return nil, fmt.Errorf("diagram must have 2 rows of equal length")
	}
	// copy the diagram into the garden
	g.row1 = []rune(rows[1])
	g.row2 = []rune(rows[2])

	return &g, nil
}

func (g *Garden) Plants(child string) (plants []string, ok bool) {
	i, ok := g.findChildIndex(child)
	if ok {
		shiftedIndex := i * 2
		plants = append(plants, rune2plant(g.row1[shiftedIndex]))
		plants = append(plants, rune2plant(g.row1[shiftedIndex+1]))
		plants = append(plants, rune2plant(g.row2[shiftedIndex]))
		plants = append(plants, rune2plant(g.row2[shiftedIndex+1]))
	}
	return plants, ok
}

func (g *Garden) findChildIndex(child string) (int, bool) {
	for i, c := range g.children {
		if c == child {
			return i, true
		}
	}
	return 0, false
}

func rune2plant(r rune) string {
	switch r {
	case 'G':
		return "grass"
	case 'C':
		return "clover"
	case 'R':
		return "radishes"
	case 'V':
		return "violets"
	default:
		return ""
	}
}
