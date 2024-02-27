package matrix

import (
	"strconv"
	"strings"
)

type Pair struct {
	Row int
	Col int
}

type Matrix struct {
	data [][]int
}

// New анализирует строку для создания новой матрицы.
func New(s string) (*Matrix, error) {
	var data [][]int

	// Разделение входной строки на строки.
	lines := strings.Split(s, "\n")
	for _, line := range lines {
		if line == "" {
			continue // Пропускаем пустые строки
		}
		var row []int
		// Разделение строки на числа, разделенные пробелами.
		numbers := strings.Split(line, " ")
		for _, numberStr := range numbers {
			// Преобразование строки в число.
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				// Возвращаем ошибку, если не удалось преобразовать строку в число.
				return nil, err
			}
			row = append(row, number)
		}
		data = append(data, row)
	}
	return &Matrix{data: data}, nil
}

func (m *Matrix) Saddle() []Pair {
	var saddlePoints []Pair

	for i, row := range m.data {
		for j, val := range row {
			if isTallestInRow(val, row) && isShortestInColumn(val, j, m.data) {
				saddlePoints = append(saddlePoints, Pair{Row: i + 1, Col: j + 1})
			}
		}
	}

	return saddlePoints
}

func isTallestInRow(value int, row []int) bool {
	for _, v := range row {
		if value < v {
			return false
		}
	}
	return true
}

func isShortestInColumn(value, columnIndex int, data [][]int) bool {
	for _, row := range data {
		if value > row[columnIndex] {
			return false
		}
	}
	return true
}
