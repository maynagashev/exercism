/*
Package accumulate contains the `accumulate` function, which, given a collection and an
operation to perform on each element of the collection, returns a new
collection containing the result of applying that operation to each element of
the input collection.
*/
package accumulate

// Accumulate applies specified converter function to collections of string items
func Accumulate(items []string, converter func(string) string) []string {
	for pos, item := range items {
		items[pos] = converter(item)
	}
	return items
}
