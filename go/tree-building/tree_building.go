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

	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	for i, r := range records {
		if r.ID != i || (r.ID == r.Parent && r.ID != 0) || r.Parent > r.ID || r.Parent < 0 {
			return nil, errors.New("invalid record")
		}
		m[r.ID] = &Node{ID: r.ID}
		if r.ID != r.Parent {
			m[r.Parent].Children = append(m[r.Parent].Children, m[r.ID])
		}
	}
	return m[0], nil
}
