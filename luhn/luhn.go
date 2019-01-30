// Package luhn applies the luhn algorithm to identification sequences
package luhn

import (
	"regexp"
	"strings"
	"unicode"
)

// Valid applies the luhn formula to a string to determine
//	if it is a valid identification number, which follows
//	the luhn algorithm.
func Valid(s string) bool {
	// replace space characters
	re := regexp.MustCompile("\\s")
	s = re.ReplaceAllLiteralString(reverse(s), "")
	// invalidate with alphabet char present
	re = regexp.MustCompile("\\D")
	if re.FindAllString(s, -1) != nil {
		return false
	}
	// catch and return insufficient length strings
	if len(s) < 2 {
		return false
	}

	// remove any characters that are non numbers that
	//	dont constitue an invalidation above
	n := strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsNumber(r) && unicode.IsLetter(r)
	})
	// instantiate slice to hold ints to be summed
	var ns = []int{}
	// iterate through our transformed string of ints and
	//	apply the doubling rule.
	for i, r := range n {
		if (i+1)%2 == 0 {
			if d := charToNum(r) * 2; d > 9 {
				d -= 9
				ns = append(ns, d)
			} else {
				ns = append(ns, d)
			}
			continue
		}
		ns = append(ns, charToNum(r))
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

// reverse 's the direction of a string
func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}
