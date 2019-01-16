package scale

import (
	"fmt"
	"strings"
)

// Global declarations for ease of process in function
var baseNotes = []string{"A", "B", "C", "D", "E", "F", "G"}
var halfStepIntervalNotes = []string{"B", "E"}

// Scale acceps a tonic and constructs a scale based on
//	a provided series of either half or whole steps
//	in intervals
func Scale(tonic string, intervals string) (scale []string) {
	scale = []string{tonic}
	// Initialize intervals as chromatic scale
	if intervals == "" {
		intervals = "mmmmmmmmmmmm"
	}
	interval := strings.Split(intervals, "")
	for _, value := range interval[:len(interval)-1] {
		fmt.Printf("-------------------------------\n%v\n-------------------------------\n", scale)
		scale = constructStep(scale[:], value)
	}
	return scale
}

// constructStep accepts an array of string notes, an interval,
//	and returns a note one indicated step above.
func constructStep(scale []string, step string) (newScale []string) {
	newScale = scale[:]
	isFlat, isSharp, isHalfStepIntervalNote := getNoteMeta(scale)
	baseNote := strings.Replace(strings.Replace(newScale[len(newScale)-1], "#", "", -1), "b", "", -1)
	nextNote := nextBaseNote(baseNote)
	// fmt.Printf("isFlat:%v\nisSharp:%v\nisHalfStep:%v\nstep:%v\nnextNote:%v\n\n", isFlat, isSharp, isHalfStepIntervalNote, step, nextNote)
	if isHalfStepIntervalNote {
		if step == "M" {
			if isFlat {
				newScale = append(newScale, nextNote)
				return newScale
			}
			if isSharp {
				newScale = append(newScale, nextBaseNote(nextNote))
				return newScale
			}
			newScale = append(newScale, nextNote+"#")
			return newScale
		} else if step == "m" {
			if isFlat {
				newScale = append(newScale, nextNote+"b")
				return newScale
			}
			if isSharp {
				newScale = append(newScale, nextNote+"#")
				return newScale
			}
			newScale = append(newScale, nextNote)
			return newScale
		}
	} else {
		if step == "M" {
			if isFlat {
				newScale = append(newScale, nextNote+"b")
				return newScale
			}
			if isSharp {
				newScale = append(newScale, nextNote+"#")
				return newScale
			}
			newScale = append(newScale, nextNote)
			return newScale
		} else if step == "m" {
			if isFlat {
				newScale = append(newScale, strings.Split(newScale[len(newScale)-1], "")[0])
				return newScale
			}
			if isSharp {
				newScale = append(newScale, nextNote)
				return newScale
			}
			newScale = append(newScale, baseNote+"#")
			return newScale
		}
	}
	return newScale
}

// getNoteMeta will return if a note is flat, and is a half step interval note
func getNoteMeta(scale []string) (isFlat bool, isSharp bool, isHalfStepIntervalNote bool) {
	fmt.Printf("isFlat?: ")
	if valueInArray("b", strings.Split(scale[len(scale)-1], "")) {
		isFlat = true
	}
	fmt.Printf("isSharp?: ")
	if valueInArray("#", strings.Split(scale[len(scale)-1], "")) {
		isSharp = true
	}
	fmt.Printf("isHalfStep?: ")
	if valueInArray(scale[len(scale)-1], halfStepIntervalNotes) {
		isHalfStepIntervalNote = true
	}
	return isFlat, isSharp, isHalfStepIntervalNote
}

// valueInArray checks if an array of string contains a string
func valueInArray(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			fmt.Printf("%v in %v\n", a, b)
			return true
		}
	}
	fmt.Printf("%v not in %v\n", a, list)
	return false
}

// nextBaseNote returns the next alphabetical note
func nextBaseNote(note string) (next string) {
	for i, value := range baseNotes {
		if value == note {
			if i < len(baseNotes)-1 {
				next = baseNotes[i+1]
				break
			}
			next = baseNotes[0]
			break
		}
	}
	return next
}
