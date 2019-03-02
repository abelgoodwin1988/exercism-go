// Package isbn ...
package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN ...
func IsValidISBN(candidate string) bool {
	// Validate properly formatted isbn
	candidate = strings.Replace(candidate, "-", "", -1)
	if len(candidate) != 10 {
		return false
	}
	var sum int
	for i, r := range candidate {
		var c int
		// catch non-digit, non-check characters
		if !unicode.IsDigit(r) && i != 9 {
			return false
		}
		// evaluate check-character position and perform
		//	summing
		if !unicode.IsDigit(r) && i == 9 {
			if strings.ToUpper(string(r)) == "X" {
				c = 10
			} else {
				return false
			}
		} else {
			c = int(r - '0')
		}
		sum += c * (10 - i)
	}
	return sum%11 == 0
}
