// Package pangram contains method to determine is a word is a panagram
package pangram

import "strings"

var alphabet = map[rune]bool{
	'a': false,
	'b': false,
	'c': false,
	'd': false,
	'e': false,
	'f': false,
	'g': false,
	'h': false,
	'i': false,
	'j': false,
	'k': false,
	'l': false,
	'm': false,
	'n': false,
	'o': false,
	'p': false,
	'q': false,
	'r': false,
	's': false,
	't': false,
	'u': false,
	'v': false,
	'w': false,
	'x': false,
	'y': false,
	'z': false,
}

// IsPangram accepts and strings and returns a boolean
//	for if it is a pangram.
func IsPangram(s string) bool {
	// get a copy of alphabet
	newAlphabet := make(map[rune]bool)
	for key, val := range alphabet {
		newAlphabet[key] = val
	}
	// set present values to true
	for _, val := range strings.ToLower(s) {
		newAlphabet[val] = true
	}
	// check for any false key values (missing alphabet)
	for _, val := range newAlphabet {
		if val == false {
			return false
		}
	}
	return true
}
