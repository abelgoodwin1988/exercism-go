// Package etl performs a data migration from old scrabble data to new scrabble
//	data format for easier processing/lookup/use
package etl

import "strings"

// Transform accepts an old data set of scrabble letters and point value.
//	this data is then transformed into a new set of data and returned
func Transform(data map[int][]string) map[string]int {
	transform := map[string]int{}
	for key, value := range data {
		for _, v := range value {
			transform[strings.ToLower(v)] = key
		}
	}
	return transform
}
