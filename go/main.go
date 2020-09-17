package main

import (
	"fmt"
	"tree-building"
)

var input = []tree.Record{
	{ID: 0},
	{ID: 1, Parent: 0},
	{ID: 2, Parent: 0},
}

func main() {
	a, _ := tree.Build(input)

	fmt.Printf("%#q", a)

}
