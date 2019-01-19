// Package strange has function ToRNA that accepts a string
//	representing a DNA strand and returns an RNA copy
package strand

// create a type to hold mapping for dna complements
type nucleotideComplements map[rune]string

// load complements with nucelotide complement values
var complements = nucleotideComplements{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA accepts a string representing DNA
//	and returns a copy of it's complementary
//	nucleoties in a string.
// Error's are not handled in this function, as
//	such, only valid dna strands are accepted
func ToRNA(dna string) (rna string) {
	for _, value := range dna {
		rna += complements[value]
	}
	return rna
}
