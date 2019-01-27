// Package protein ...
package protein

import "errors"

var codonToAmino = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

// ErrStop is an exported default type error used
//	when a stop codon is encountered.
var ErrStop = errors.New("Stop Codon")

// ErrInvalidBase is an exportted default type error used
//	when a non-existant dna codon is traversed.
var ErrInvalidBase = errors.New("Invalid Codon")

// FromRNA takes a codon and decodes it to an amino acid
func FromRNA(rna string) (amino []string, err error) {
	rnaSlice := splitByCharCount(rna, 3)
	for _, codon := range rnaSlice {
		if protein, err := FromCodon(codon); err == nil {
			amino = append(amino, protein)
		} else if err.Error() == "Stop Codon" {
			return amino, nil
		} else {
			return amino, ErrInvalidBase
		}
	}
	return amino, nil
}

// FromCodon takes a codon rna seq and returns the approrpiate
//	protein
func FromCodon(codon string) (protein string, err error) {
	if val, ok := codonToAmino[codon]; ok {
		if val == "STOP" {
			return "", ErrStop
		}
		return val, nil
	}
	return "", ErrInvalidBase
}

func splitByCharCount(word string, count int) []string {
	var strSlice = []string{}
	iterations := int(len(word) / count)
	for i := 0; i < iterations; i++ {
		strSlice = append(strSlice, word[i*count:i*count+count])
	}
	return strSlice
}
