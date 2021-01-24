package sublist

// Relation describes return type
type Relation string

// Sublist determines if the first list is contained within the second
//list, if the second list is contained within the first list, if both lists are
//contained within each other or if none of these are true
func Sublist(a []int, b []int) Relation {
	switch {
	case len(a) == len(b) && isSublist(b, a):
		return "equal"
	case len(a) > len(b) && isSublist(b, a):
		return "superlist"
	case len(a) < len(b) && isSublist(a, b):
		return "sublist"
	default:
		return "unequal"
	}
}

func isSublist(small []int, big []int) bool {
	for j := 0; j < len(big)-len(small)+1; j++ {
		match := true
		for i, n := range small {
			if n != big[i+j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}
