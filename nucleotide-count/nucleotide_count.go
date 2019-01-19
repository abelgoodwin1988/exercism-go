package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, value := range strings.Split(string(d), "") {
		var nonNucleotide bool
		for nucleotide := range h {
			n := string(nucleotide)
			if value == n {
				h[nucleotide]++
				nonNucleotide = false
				break
			} else {
				nonNucleotide = true
			}
		}
		if nonNucleotide {
			return h, errors.New("Non Nucleotide in strand")
		}
	}
	return h, nil
}
