/*
Package tree contains Build method that implements logic for building tree structure from unordered slice of records.
*/
package tree

import (
	"errors"
	"sort"
)

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

// Build builds tree of Node-s from slice of records.
// Constraints:
//    1. The ID number is always between 0 (inclusive) and the length of the record list (exclusive).
//    2. All records have a parent ID lower than their own ID
//    3. Root record has a parent ID that's equal to its own ID.
func Build(records []Record) (*Node, error) {

	m := make(map[int]*Node, len(records))

	// sorting, possible root will be bubbled to the first position
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	rootID := 0
	for i, r := range records {
		if r.ID != i || (r.ID == r.Parent && r.ID != rootID) || r.ID < r.Parent {
			return nil, errors.New("invalid record")
		}
		m[r.ID] = &Node{ID: r.ID}
		parent, ok := m[r.Parent]
		if !ok {
			return nil, errors.New("invalid parent")
		}
		if r.ID != r.Parent {
			parent.Children = append(parent.Children, m[r.ID])
		}
	}
	return m[rootID], nil
}
