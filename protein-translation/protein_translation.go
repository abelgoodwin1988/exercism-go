// Package protein ...
package protein

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
var aminoToCodon = map[string][]string{
	"Methionine":    {"AUG"},
	"Phenylalanine": {"UUU", "UUC"},
	"Leucine":       {"UUA", "UUG"},
	"Serine":        {"UCU", "UCG", "UCA", "UCC"},
	"Tyrosine":      {"UAU", "UAC"},
	"Cysteine":      {"UGU", "UGC"},
	"Tryptophan":    {"UGG"},
	"STOP":          {"UAA", "UAG", "UGA"},
}

// FromRNA takes a codon and decodes it to an amino acid
func FromRNA(rna string) (amino string, err ErrStop) {
	rnaSlice := splitByCharCount(rna, 3)
	print(rnaSlice)
	return amino, err
}

// FromCodon takes a codon rna seq and returns the approrpiate
//	protein
func FromCodon(codon string) (protein string, err error) {

	return protein, err
}

func splitByCharCount(word string, count int) []string {
	var strSlice = []string{}
	iterations := int(len(word)/count) + 1
	for i := 0; i <= iterations; i++ {
		strSlice = append(strSlice, word[i*count:i*count+count])
	}
	return []string{""}
}
