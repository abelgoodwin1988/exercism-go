package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides.
type DNA []rune

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
	// Iterate through the string as an array of runes (runes to match
	//	histogram mapping). Attempt to access the hisogram map
	//	at the specified value location, and if successful, that
	//	nucleotide exists, if err we have a non-nucleotide and
	//	should return an err immediately.
	for _, value := range d {
		if _, ok := h[value]; !ok {
			return h, fmt.Errorf("bad nucleotide %s", string(value))
		} else {
			h[value]++
		}
	}
	return h, nil
}
