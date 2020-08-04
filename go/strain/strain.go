package strain

// Ints is collection of integers
type Ints []int

// Lists is collection of integer slices
type Lists [][]int

// Strings is collection of strings
type Strings []string

// Keep filters receiver collection by given predicate function
func (col Ints) Keep(pred func(int) bool) (res Ints) {
	for _, v := range col {
		if pred(v) {
			res = append(res, v)
		}
	}
	return
}

// Discard removes elements from receiver collection by given predicate function
func (col Ints) Discard(pred func(int) bool) (res Ints) {
	for _, v := range col {
		if !pred(v) {
			res = append(res, v)
		}
	}
	return

}

// Keep filters receiver collection by given predicate function
func (col Lists) Keep(pred func([]int) bool) (res Lists) {
	for _, v := range col {
		if pred(v) {
			res = append(res, v)
		}
	}
	return
}

// Keep filters receiver collection by given predicate function
func (col Strings) Keep(pred func(string) bool) (res Strings) {
	for _, v := range col {
		if pred(v) {
			res = append(res, v)
		}
	}
	return
}
