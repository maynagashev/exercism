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

	if len(records) == 0 {
		return nil, nil
	}

	return &Node{}, nil
}
