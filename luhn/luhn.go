// Package luhn applies the luhn algorithm to identification sequences
package luhn

import (
	"unicode"
)

// Valid applies the luhn formula to a string to determine
//	if it is a valid identification number, which follows
//	the luhn algorithm.
func Valid(s string) bool {
	/**
	 * Clean & Validate string
	 */
	// Replace space, if non-number invalidate.
	cleanedString := ""
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}
		if string(r) == "-" {
			return false
		}
		if unicode.IsLetter(r) {
			return false
		}
		if unicode.IsSymbol(r) {
			return false
		}
		cleanedString += string(r)
	}
	// catch and return insufficient length strings
	if len(cleanedString) < 2 {
		return false
	}

	// instantiate slice to hold ints to be summed
	var ns = []int{}
	// iterate through our transformed string of ints and
	//	apply the doubling rule.
	var count int
	for i := len(cleanedString) - 1; i >= 0; i-- {
		r := []rune(cleanedString)[i]
		if (count+1)%2 == 0 {
			if d := charToNum(r) * 2; d > 9 {
				d -= 9
				ns = append(ns, d)
			} else {
				ns = append(ns, d)
			}
			count++
			continue
		}
		ns = append(ns, charToNum(r))
		count++
		continue
	}
	// sum the sequence of numbers
	var sum int
	for _, val := range ns {
		sum += val
	}
	// if divisible by 10, it's valid!
	return (sum%10 == 0)
}

// charToNum converts a rune char in to an int
func charToNum(r rune) int {
	if '0' <= r && r <= '9' {
		return int(r) - '0'
	}
	return 0
}
