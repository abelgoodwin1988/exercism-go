// Package encode enables RLE encoding
package encode

import (
	"strconv"
	"unicode"
)

// RunLengthDecode takes an encoded RLE and decodes it
//	to a lossless string
func RunLengthDecode(s string) string {
	num := ""
	stringer := ""
	for _, r := range s {
		if unicode.IsNumber(r) {
			if num == "" {
				num = string(r)
				continue
			} else {
				num += string(r)
			}
		} else {
			if val, ok := strconv.Atoi(num); ok == nil {
				for i := 0; i < val; i++ {
					stringer += string(r)
				}
			} else {
				stringer += string(r)
			}
			num = ""
		}
	}
	return stringer
}

// RunLengthEncode takes a normal string of spaces/alphabets
//	and encodes it by counting a continuous series of chars
func RunLengthEncode(s string) string {
	// If the string is empty, or one character,
	//	simply return the string
	if len(s) < 2 {
		return s
	}
	// iterators
	j := 1
	stringer := ""
	// iterate through the string
	for i, v := range s {
		// Lookahead match, if end of list ...
		if i != len(s)-1 {
			// Compare lookahead
			if v == rune(s[i+1]) {
				j++
				continue
			} else {
				if j > 1 {
					stringer += strconv.Itoa(j) + string(v)
					j = 1
					continue
				} else {
					stringer += string(v)
					j = 1
					continue
				}
			}
		} else {
			if j > 1 {
				stringer += strconv.Itoa(j) + string(v)
			} else {
				stringer += string(v)
			}
		}
	}
	return stringer
}
