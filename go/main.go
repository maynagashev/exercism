package main

import (
	"encoding/json"
	"fmt"
	"local/exercism/tournament"
	"os"
	"strings"
)

func main() {
	input := strings.NewReader(`
Courageous Californians;Devastating Donkeys;win
Allegoric Alaskians;Blithering Badgers;win
Devastating Donkeys;Allegoric Alaskians;loss
Courageous Californians;Blithering Badgers;win
Blithering Badgers;Devastating Donkeys;draw
Allegoric Alaskians;Courageous Californians;draw
`)
	tournament.Tally(input, os.Stdout)
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
