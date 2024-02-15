package linkedlist

import (
	"errors"
)

// Define List and Node types here.
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

type List struct {
	head *Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.

// NewList returns a new list with all the elements passed as arguments
func NewList(elements ...interface{}) *List {
	l := &List{}
	for _, e := range elements {
		l.Push(e)
	}
	return l
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

// Unshift adds a new node to the beginning of the list
func (l *List) Unshift(v interface{}) {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n
	} else {
		n.next = l.head
		l.head.prev = n
		l.head = n
	}
}

// Push adds a new node to the end of the list
func (l *List) Push(v interface{}) {
	n := &Node{Value: v}
	if l.head == nil {
		l.head = n
	} else {
		current := l.Last()
		current.next = n
		n.prev = current
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("empty list")
	}
	v := l.head.Value
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	return v, nil
}

// Pop removes the last node from the list and returns its value
func (l *List) Pop() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("empty list")
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	if current.prev == nil {
		l.head = nil
	} else {
		current.prev.next = nil
	}
	return current.Value, nil
}

func (l *List) Reverse() {
	current := l.head
	var prev *Node
	var next *Node
	for current != nil {
		next = current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
	l.head = prev
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	if l.head == nil {
		return nil
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	return current
}
