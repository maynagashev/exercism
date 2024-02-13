package main

import (
	"encoding/json"
	"fmt"
	"local/exercism/tournament"
	"os"
	"strings"
)

var str = `This is a raw string literal.
It can contain special characters such as \n, \t, ", ', and even ` + "`" + ` backticks.
It doesn't interpret escape sequences.`

func main() {
	input := strings.NewReader(`
Courageous Californians;Devastating Donkeys;win
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;loss
Courageous Californians;Blithering Badgers;win
Blithering Badgers;Devastating Donkeys;draw
Allegoric Alaskians;Courageous Californians;draw
`)
	err := tournament.Tally(input, os.Stdout)
	if err != nil {
		fmt.Println(err)
	}
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
