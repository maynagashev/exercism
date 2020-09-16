/*
Package tree contains Build method that implements logic for building tree structure from slice of records.
*/
package tree

// Record struct representing single record
type Record struct {
	ID     int
	Parent int
}

// Node struct representing result tree
type Node struct {
	ID       int
	Children []*Node
}

// Build builds tree of Node-s from slice of records
func Build(records []Record) (*Node, error) {
	m := make(map[int]*Node, len(records))
	var rootID int

	if len(records) == 0 {
		return nil, nil
	}
	// making map
	for _, r := range records {
		m[r.ID] = &Node{r.ID, []*Node{}}
	}

	for _, r := range records {
		if r.ID == r.Parent {
			rootID = r.ID
		}

		node := m[r.ID]
		parent := m[r.Parent]

		parent.Children = append(parent.Children, node)
	}

	return m[rootID], nil
}
