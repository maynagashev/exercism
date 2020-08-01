/*
Package strand contains function for DNA -> RNA strand conversion
*/
package strand

import (
	"strings"
)

var toRNA = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

// ToRNA transcribes DNA strand to RNA strand
func ToRNA(dna string) (rna string) {
	for _, s := range strings.Split(dna, "") {
		rna += toRNA[s]
	}
	return
}
