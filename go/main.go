package main

import (
	"encoding/json"
	"fmt"
	"local/exercism/tree-building"
)

var input = []tree.Record{
	{ID: 0, Parent: 0},
	{ID: 1, Parent: 0},
	{ID: 2, Parent: -1},
}

func main() {
	res, err := tree.Build(input)
	//fmt.Printf("%+v \t %+v \n", input, m)
	//fmt.Printf("%#v \t %#v \n", input, m)
	PrettyPrint(input)
	fmt.Println("==Result==")
	PrettyPrint(res)
	fmt.Println("==Err==")
	fmt.Println(err)
	//PrettyPrint(m)
}

func PrettyPrint(data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}
