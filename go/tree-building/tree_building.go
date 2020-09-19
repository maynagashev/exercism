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

// Build builds tree of Node-s from slice of records
func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	// making map of Node instances
	m := make(map[int]*Node, len(records))
	for _, r := range records {
		if _, isAlreadyMapped := m[r.ID]; isAlreadyMapped {
			return nil, errors.New("duplicated node id")
		}
		m[r.ID] = &Node{r.ID, nil}
	}

	// sorting, possible root will bubbled to the first position
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
		node := m[r.ID]
		parent, isParentExists := m[r.Parent]
		if !isParentExists {
			return nil, errors.New("invalid parent")
		}
		if r.ID != r.Parent {
			parent.Children = append(parent.Children, node)
		}
		lastID = r.ID
	}
	return m[rootID], nil
}
