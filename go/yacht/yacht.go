package yacht

// Score calculates the score of a dice roll for a given Yacht category.
func Score(dice []int, category string) int {
	counts := make(map[int]int)
	total := 0
	// Count the number of each dice value and calculate the total
	for _, value := range dice {
		counts[value]++
		total += value
	}

	switch category {
	case "ones":
		return counts[1] * 1
	case "twos":
		return counts[2] * 2
	case "threes":
		return counts[3] * 3
	case "fours":
		return counts[4] * 4
	case "fives":
		return counts[5] * 5
	case "sixes":
		return counts[6] * 6
	case "full house":
		twoFound, threeFound := false, false
		for _, count := range counts {
			if count == 2 {
				twoFound = true
			} else if count == 3 {
				threeFound = true
			}
		}
		if twoFound && threeFound {
			return total
		}
	case "four of a kind":
		for num, count := range counts {
			if count >= 4 {
				return num * 4
			}
		}
	case "little straight":
		if len(counts) == 5 && counts[1] == 1 && counts[5] == 1 {
			return 30
		}
	case "big straight":
		if len(counts) == 5 && counts[2] == 1 && counts[6] == 1 {
			return 30
		}
	case "choice":
		return total
	case "yacht":
		for _, count := range counts {
			if count == 5 {
				return 50
			}
		}
	}
	return 0
}
