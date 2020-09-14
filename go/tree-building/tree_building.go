package tree

type Record struct {
	ID int
	Parent int
}

type Node struct {
	ID int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	return &Node{}, nil
}