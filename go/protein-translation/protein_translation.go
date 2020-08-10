// Package protein contains translation methods for RNA sequences
package protein

import "errors"

var (
	// ErrStop error when there is STOP codon in a sequence
	ErrStop = errors.New("stop codon")
	// ErrInvalidBase error when codon is invalid and not found in base
	ErrInvalidBase = errors.New("invalid codon base")
)

// FromCodon translates codon to protein
func FromCodon(codon string) (p string, err error) {
	switch codon {
	case "AUG":
		p = "Methionine"
	case "UUU", "UUC":
		p = "Phenylalanine"
	case "UUA", "UUG":
		p = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		p = "Serine"
	case "UAU", "UAC":
		p = "Tyrosine"
	case "UGU", "UGC":
		p = "Cysteine"
	case "UGG":
		p = "Tryptophan"
	case "UAA", "UAG", "UGA": // STOP
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
	return p, nil
}

// FromRNA translates RNA sequence to protein sequence
func FromRNA(rna string) (proteins []string, err error) {
	for i := 0; i < len(rna); i += 3 {
		p, err := FromCodon(rna[i : i+3])
		if err != nil {
			if err == ErrStop {
				break
			}
			return proteins, err
		}
		proteins = append(proteins, p)
	}
	return proteins, nil
}
