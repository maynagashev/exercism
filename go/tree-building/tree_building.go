/*
Package tree contains Build method that implements logic for building tree structure from slice of records.
*/
package tree

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
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
	m := make(map[int]*Node, len(records))
	rootID := -1

	if len(records) == 0 {
		return nil, nil
	}

	// making map
	for _, r := range records {
		if _, isAlreadyMapped := m[r.ID]; isAlreadyMapped {
			return nil, errors.New("duplicated node id")
		}
		m[r.ID] = &Node{r.ID, nil}

	}

	// sorting
	sort.Slice(records, func(i, j int) bool { return records[i].ID < records[j].ID })

	// linking parents
	for _, r := range records {
		if r.ID == r.Parent {
			rootID = r.ID
		}
		if r.ID<r.Parent {
			return nil, errors.New("parent younger than child, might be possible circular dependency")
		}

		node := m[r.ID]
		parent, isParentExists := m[r.Parent]
		if !isParentExists {
			return nil, errors.New("invalid parent")
		}

		if r.ID != r.Parent {
			parent.Children = append(parent.Children, node)
		}
	}
	fmt.Printf("root=%s\n", strconv.Itoa(rootID))
	fmt.Printf("map=%v\n", m)

	root, ok := m[rootID]
	if !ok {
		return nil, errors.New("no root")
	}

	return root, nil
}
