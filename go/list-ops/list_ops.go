package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Foldl fold (reduce) each item into the accumulator from the left
func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _, v := range s {
		initial = fn(initial, v)
	}
	return initial
}

// Foldr fold (reduce) each item into the accumulator from the right
func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	for i := len(s) - 1; i >= 0; i-- {
		initial = fn(s[i], initial)
	}
	return initial
}

// Filter given a predicate and a list, return the list of all items for which predicate(item) is True
func (s IntList) Filter(fn func(int) bool) IntList {
	for i := 0; i < len(s); i++ {
		if !fn(s[i]) {
			s = append(s[:i], s[i+1:]...) // remove the element
			i--
		}
	}
	return s
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	for i := 0; i < len(s); i++ {
		s[i] = fn(s[i])
	}
	return s
}

func (s IntList) Reverse() IntList {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	for _, l := range lists {
		s = append(s, l...)
	}
	return s
}
