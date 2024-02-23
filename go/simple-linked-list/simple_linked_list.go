package linkedlist

import "errors"

// Define the List and Element types here.
type Element struct {
	Value int
	Next  *Element
}

type List struct {
	head *Element
	size int
}

func New(elements []int) *List {
	list := &List{}
	for _, element := range elements {
		list.Push(element)
	}
	return list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	newEl := &Element{Value: element}
	if l.head == nil {
		l.head = newEl
	} else {
		current := l.head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newEl
	}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("empty list")
	}
	current := l.head
	if current.Next == nil {
		l.head = nil
		l.size--
		return current.Value, nil
	}
	for current.Next.Next != nil {
		current = current.Next
	}
	last := current.Next
	current.Next = nil
	l.size--
	return last.Value, nil
}

func (l *List) Array() []int {
	var array []int
	current := l.head
	for current != nil {
		array = append(array, current.Value)
		current = current.Next
	}
	return array
}

func (l *List) Reverse() *List {
	var reversed []int
	current := l.head
	for current != nil {
		// Prepend the current value to the reversed list
		reversed = append([]int{current.Value}, reversed...)
		current = current.Next
	}
	return New(reversed)
}
