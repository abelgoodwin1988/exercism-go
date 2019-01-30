// Package wordcount count words
package wordcount

import (
	"strconv"
	"strings"
	"unicode"
)

// Frequency is an histogram of the words in a phrase
type Frequency map[string]int

// WordCount counts the word in a sentence
func WordCount(input string) Frequency {
	var frequencyCount Frequency = map[string]int{}
	for _, word := range strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != 39
	}) {
		word := strings.ToLower(word)
		newWord, err := strconv.Unquote(strings.Replace(word, "'", "\"", -1))
		if err == nil {
			frequencyCount[newWord]++
			continue
		}
		frequencyCount[word]++
	}

	return frequencyCount
}
