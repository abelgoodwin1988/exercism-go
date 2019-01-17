package hamming

import (
	"errors"
	"strings"
)

// Distance calculates hamming difference between two
//	DNA strands; that is to say the difference from
//	each nucleic acid in order of the DNA.
func Distance(a, b string) (int, error) {
	// catch empy-string examples
	if a == "" && b == "" {
		return 0, nil
	}
	// Split the strands into arrays for easier
	//	iteration comparison
	aSplit := strings.Split(a, "")
	bSplit := strings.Split(b, "")

	// Check for equal-length strands, err if unequal
	if len(aSplit) != len(bSplit) {
		return 0, errors.New("unequal strand lengths")
	}

	// Initalize hamming at default nil value of 0
	var hammingDifference int

	// Iterate through strand-array and compare index
	//	of each strand for inequality.
	//	When there is inequality tic a value to the
	//	hamming difference
	for i := range a {
		if aSplit[i] != bSplit[i] {
			hammingDifference++
		}
	}
	return hammingDifference, nil
}
