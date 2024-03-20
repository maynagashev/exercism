package rectangles

import (
	"errors"
)

type Point struct {
	x int
	y int
}

func Count(diagram []string) int {
	res := 0
	var rowRes int
	// Scan the diagram and for every left top corner, find the right bottom corner
	for y, row := range diagram {
		rowRes = 0
		for x, r := range row {
			if r == '+' {
				found := findAllRectanglesFromUpperLeft(Point{x, y}, diagram)
				res += found
				rowRes += found
				//fmt.Printf("\tRow %d, Col %d: %d rectangles\n", y, x, found)
			}
		}
		//fmt.Printf("Row %d: %d rectangles\n", y, rowRes)
	}
	return res
}

func findAllRectanglesFromUpperLeft(upperLeft Point, diagram []string) int {
	x := upperLeft.x
	y := upperLeft.y
	res := 0

	// Find the right bottom corner
	// Iterate over current (upper) row and find the first '+' after the current x
	for i := x + 1; i < len(diagram[y]); i++ {
		if diagram[y][i] == '+' {
			bottomRight, err := findBottomRight(i, y, diagram)
			// Iterate over the rows below the current row and find ALL the bottom corners
			for err == nil {
				if validateRectangle(upperLeft, bottomRight, diagram) {
					res++
					//fmt.Printf("\t\tFound rectangle: %v, %v\n", upperLeft, bottomRight)
				}
				// search for the next bottom corner from the current bottom corner
				bottomRight, err = findBottomRight(i, bottomRight.y, diagram)
			}
		}
	}

	return res
}

func findBottomRight(x int, y int, diagram []string) (Point, error) {
	for j := y + 1; j < len(diagram); j++ {
		if diagram[j][x] == '+' {
			return Point{x, j}, nil
		}
	}
	return Point{}, errors.New("no bottom corner found")
}

func validateRectangle(upperLeft Point, bottomRight Point, diagram []string) bool {
	// check top and bottom row
	for x := upperLeft.x + 1; x < bottomRight.x; x++ {
		if !isHorizontal(diagram[upperLeft.y][x]) || !isHorizontal(diagram[bottomRight.y][x]) {
			return false
		}
	}

	// check left and right column
	for y := upperLeft.y + 1; y < bottomRight.y; y++ {
		if !isVertical(diagram[y][upperLeft.x]) || !isVertical(diagram[y][bottomRight.x]) {
			return false
		}
	}

	// check left bottom corner, because before we only checked the right bottom corner
	if diagram[bottomRight.y][upperLeft.x] != '+' {
		return false
	}

	return trueexercism download --track=go --exercise=grade-school

}

func isVertical(u uint8) bool {
	return u == '|' || u == '+'
}

func isHorizontal(u uint8) bool {
	return u == '-' || u == '+'
}
