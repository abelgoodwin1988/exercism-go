// Package isogram allows IsIsogram to be accessed by same-package
//	test suite
package isogram

import "unicode"

// IsIsogram accepts characters and determines if that set of characters
// is an isogram and isogram is a word with no repeating characters;
//	however spaces
func IsIsogram(characters string) bool {
	var charSlice = []rune{}
	for _, value := range characters {
		value = unicode.ToLower(value)
		if isInSlice(value, charSlice) {
			return false
		}
		charSlice = append(charSlice, value)
	}
	return true
}

func isInSlice(value rune, slice []rune) bool {
	for _, r := range slice {
		if value == r && value != '-' && value != ' ' {
			return true
		}
	}
	return false
}
