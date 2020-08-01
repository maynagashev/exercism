package main

import (
	"fmt"
	"strings"
)

var toRNA = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

func main() {
	ToRNA("GCTA")
}

func ToRNA(dna string) (rna string) {
	for _, s := range strings.Split(dna, "") {
		rna += toRNA[s]
	}
	fmt.Println(rna)
	return
}
