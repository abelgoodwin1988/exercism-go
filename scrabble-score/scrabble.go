// Package scrabble creates functionality for the scrabble game wherein
//	spelling out a word gives a user different scores based on
//	the aggregate of the words point values
package scrabble

import "strings"

// pointValues places letters into index of their point value
var pointValues = [][]rune{
	{},
	{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'},
	{'D', 'G'},
	{'B', 'C', 'M', 'P'},
	{'F', 'H', 'V', 'W', 'Y'},
	{'K'},
	{},
	{},
	{'J', 'X'},
	{},
	{'Q', 'Z'},
}

// Score takes a work, parses as an array of rune and
//	returns a score-value for that word
func Score(word string) (score int) {
	// iterate through each rune of the provided word
	// 	iterate through the outerArray of the point-values global var
	// 		iterate through the inner array runes & equality check
	//		if equal, apply index of outer rune iteration to score
	for _, value := range strings.ToUpper(word) {
		for i, outer := range pointValues {
			for _, innerValue := range outer {
				if value == innerValue {
					score += i
				}
			}
		}
	}
	return score
}
