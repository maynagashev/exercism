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
// - The ID number is always between 0 (inclusive) and the length of the record list (exclusive).
// - All records have a parent ID lower than their own ID
// - Root record has a parent ID that's equal to its own ID (root does not have to have an ID = 0).
func Build(records []Record) (*Node, error) {

	m := make(map[int]*Node, len(records))

	// sorting, possible root will be bubbled to the first position
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	rootID, lastID := -1, -1
	for i, r := range records {
		if i == 0 {
			rootID = r.ID
		}
		if r.ID == r.Parent && r.ID != rootID {
			return nil, errors.New("cycle directly")
		}
		if r.ID < r.Parent {
			return nil, errors.New("higher id parent of lower id (possible indirect cycle)")
		}
		if lastID != -1 && r.ID != lastID+1 {
			return nil, errors.New("non-continuous")
		}
		m[r.ID] = &Node{ID: r.ID}
		parent, ok := m[r.Parent]
		if !ok {
			return nil, errors.New("invalid parent")
		}
		if r.ID != r.Parent {
			parent.Children = append(parent.Children, m[r.ID])
		}
		lastID = r.ID
	}
	return m[rootID], nil
}
