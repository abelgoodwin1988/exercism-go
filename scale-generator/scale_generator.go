// Package scale implements scale construction in function scale
//	when provided a interval-set string
package scale

import "strings"

var allNotes = [][]string{{"G#", "Ab"}, {"A"}, {"A#", "Bb"}, {"B", "Cb"}, {"B#", "C"}, {"C#", "Db"}, {"D"}, {"D#", "Eb"}, {"E", "Fb"}, {"E#", "F"}, {"F#", "Gb"}, {"G"}}

// We will be receiving a string of intervals m and M, use an int
//	to representing how much of the allNotes array woudl be
//	traversed in getting to that next note based on the interval
var steps = map[string]int{
	"m": 1,
	"M": 2,
}

// Scale iterates through intervals and gets the next note
//	the appropriate step ahead and constructs a scale as it
//	goes
func Scale(tonic string, intervals string) (scale []string) {
	// set root note of scale, which is provided tonic note
	scale = []string{tonic}
	// initialize empty interval as chromatic
	if intervals == "" {
		intervals = "mmmmmmmmmmm"
	}
	// turn intervals into an array
	intervalsArray := strings.Split(intervals, "")

	index := 0
	// endex := len(intervalsArray) - 1

	// Find the index of our provided tonic note
	for i, value := range allNotes {
		index = i
		if isInArray(tonic, value) {
			break
		}
	}
	// Get an ordered array of notes so we don't need to do any
	//	iterating to end and then start at the beginning
	//	while constructing our array
	var orderedNotes = [][]string{}
	for _, value := range allNotes[index:] {
		orderedNotes = append(orderedNotes, value)
	}
	for _, value := range allNotes[:index] {
		orderedNotes = append(orderedNotes, value)
	}
	println(orderedNotes)

	// Iterate through our intervals and construct a
	//	scale
	index = 0
	for _, interval := range intervalsArray {
		scale = append(scale, orderedNotes[index+steps[interval]][0])

		index += steps[interval]
	}
	return scale
}

// Is in array takes a string and array of string, iterates
//	through the array and checks for equality.
func isInArray(str string, array []string) bool {
	for _, value := range array {
		if str == value {
			return true
		}
	}
	return false
}
