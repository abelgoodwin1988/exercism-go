// Package wordcount ...
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a map that tracks the frequency
//	of a words in a phrase
type Frequency map[string]int

// WordCount accepts a string and returens a frequency
//	which defines the count of each word in that phrase
func WordCount(s string) Frequency {
	frequency := Frequency{}
	re := regexp.MustCompile("\\s|[\",]")
	words := re.Split(s, -1)
	for _, word := range words {
		re = regexp.MustCompile("[^a-zA-Z0-9_']")
		word = strings.ToLower(re.ReplaceAllLiteralString(word, ""))
		re = regexp.MustCompile("^'|'$")
		word = strings.ToLower(re.ReplaceAllLiteralString(word, ""))
		if word == "" {
			continue
		}
		frequency[word]++
	}
	return frequency
}
