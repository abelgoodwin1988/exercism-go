// Package luhn applies the luhn algorithm to identification sequences
package luhn

import (
	"strings"
	"unicode"
)

// Valid applies the luhn formula to a string to determine
//	if it is a valid identification number, which follows
//	the luhn algorithm.
func Valid(s string) bool {
	/**
	 * Clean & Validate string
	 */

	// Clean the string of unacceptable but interrupting characters
	//	such as spaces which are allowable, but interrupt the positional
	// calculation
	s = strings.Replace(s, " ", "", -1)

	// VALIDATION & PROCESS
	// Strings less than 2 in lengths are invalid
	if len(s) < 2 {
		return false
	}
	// Iterate over Luhn Candidate and invalidate candidates
	// containing non-luhn acceptable characters.
	//	Also sum as we go for the appropriately positioned characters
	var sum int
	var count = 0
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		// all characters in a luhn set should be a number
		if !unicode.IsDigit(r) {
			return false
		}
		// Get a int from the rune
		d := int(r - '0')
		// If this is a 'second' position, it should be doubled
		//	and if it is > 9 it should be subtracted by 9
		//	and added to the sum
		if count%2 == 1 {
			d *= 2
			if d < 0 || d > 9 {
				d -= 9
			}
		}
		sum += d
		count++
	}
	return (sum%10 == 0)
}
