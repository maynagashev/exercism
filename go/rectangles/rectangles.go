package rectangles

import "errors"

type Point struct {
	x int
	y int
}

func Count(diagram []string) int {
	res := 0
	// Scan the diagram and for every left top corner, find the right bottom corner
	for y, row := range diagram {
		for x, rune := range row {
			if rune == '+' {
				findAllRectanglesFromUpperLeft(Point{x, y}, diagram)
			}
		}
	}
	return res
}

func findAllRectanglesFromUpperLeft(upperLeft Point, diagram []string) {
	x := upperLeft.x
	y := upperLeft.y

	// Find the right bottom corner
	for i := x + 1; i < len(diagram[y]); i++ {
		if diagram[y][i] == '+' {
			upperRight := Point{i, y}
			bottomRight, err := findBottomRight(upperRight, i, diagram)
			if err != nil {
				continue
			}
			validateRectangle(upperLeft, bottomRight, diagram)
		}
	}
}

func findBottomRight(upperRight Point, i int, diagram []string) (Point, error) {
	for j := upperRight.y + 1; j < len(diagram); j++ {
		if diagram[j][upperRight.x] == '+' {
			return Point{i, j}, nil
		}
	}
	return Point{}, errors.New("no bottom corner found")
}

func validateRectangle(upperLeft Point, bottomRight Point, diagram []string) {
	//
}
