// Package lsproduct ...
package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct ...
func LargestSeriesProduct(digits string, span int) (int, error) {
	// catch span-longer-than-string error
	if len(digits) < span {
		return 1, errors.New("span longer than string")
	}
	if span == 0 {
		return 1, nil
	}
	if span < 0 {
		return 1, errors.New("Span must be greater than 0")
	}
	// initalize return value
	var largest int
	// assign reused loop value outside of loop
	dl := len(digits)
	// iterate over string from left -> right getting runes
	//	for each position and evaluating the 0...n sum
	//	for the max
	for i := range digits {
		if i+span > dl {
			break
		}
		sum := 1
		for _, m := range digits[i : i+span] {
			if !unicode.IsDigit(m) {
				return largest, errors.New("non-digit encountered")
			}
			sum *= int(m - '0')
		}
		if sum > largest {
			largest = sum
		}
	}
	return largest, nil
}
