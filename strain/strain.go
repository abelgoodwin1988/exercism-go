// Package strain adds two methods to "collection" data types
//	which, based on a callback function, will exclude/include
//	values from the "collection" and return that.
package strain

// Ints type to mimic other language collection
type Ints []int

// Lists type to mimic other language collection
type Lists [][]int

// Strings type to mimic other language collection
type Strings []string

// Keep (Ints) returns a slice of int filtered by a provided function
func (i Ints) Keep(keep func(int) bool) (ints Ints) {
	for _, value := range i {
		if keep(value) {
			ints = append(ints, value)
		}
	}
	return ints
}

// Discard (Ints) returns a slice of int filtered by a provided function
func (i Ints) Discard(discard func(int) bool) (ints Ints) {
	for _, value := range i {
		if !discard(value) {
			ints = append(ints, value)
		}
	}
	return ints
}

// Keep ([]Ints) returns a slice of int filtered by a provided function
func (l Lists) Keep(keep func([]int) bool) (lists Lists) {
	for _, value := range l {
		if keep(value) {
			lists = append(lists, value)
		}
	}
	return lists
}

// Keep (Strings) returns a slice of int filtered by a provided function
func (s Strings) Keep(keep func(string) bool) (strings Strings) {
	for _, value := range s {
		if keep(value) {
			strings = append(strings, value)
		}
	}
	return strings
}
