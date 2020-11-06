package pascal

// Triangle computes Pascal's triangle up to a given number of rows
func Triangle(n int) (resultRows [][]int) {
	for i := 0; i < n; i++ {
		row := []int{1}
		if i-1 < 0 {
			resultRows = append(resultRows, row)
			continue
		}
		prevRow := resultRows[i-1]
		for j := 1; j <= i; j++ {
			samePosValue := 0
			if j < len(prevRow) {
				samePosValue = prevRow[j]
			}
			row = append(row, prevRow[j-1]+samePosValue)
		}
		resultRows = append(resultRows, row)
	}
	return
}
