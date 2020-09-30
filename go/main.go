package main

import (
	"encoding/json"
	"fmt"
	"local/exercism/robot-name"
)

func main() {
	r := robotname.Robot{}
	r2 := robotname.Robot{}
	fmt.Println(r.Name())
	fmt.Println(r2.Name())
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
