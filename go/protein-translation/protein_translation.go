// Package protein contains translation methods for RNA sequences
package protein

import "errors"

var (
	// ErrStop error when there is STOP codon in a sequence
	ErrStop = errors.New("stop codon")
	// ErrInvalidBase error when codon is invalid and not found in base
	ErrInvalidBase = errors.New("invalid codon base")

	proteins = map[string][]string{
		"Methionine":    {"AUG"},
		"Phenylalanine": {"UUU", "UUC"},
		"Leucine":       {"UUA", "UUG"},
		"Serine":        {"UCU", "UCC", "UCA", "UCG"},
		"Tyrosine":      {"UAU", "UAC"},
		"Cysteine":      {"UGU", "UGC"},
		"Tryptophan":    {"UGG"},
		"STOP":          {"UAA", "UAG", "UGA"},
	}
)

// FromCodon translates codon to protein
func FromCodon(codon string) (protein string, err error) {
	for p, codons := range proteins {
		for _, c := range codons {
			if codon == c && p == "STOP" {
				return "", ErrStop
			} else if codon == c {
				return p, nil
			}
		}
	}
	return "", ErrInvalidBase
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
