// Package accumulate is given a
package accumulate

// Accumulate takes a collection and performs a function on all
//	items of that collection and return it
func Accumulate(collection []string, operation func(string) string) (newCollection []string) {
	for _, value := range collection {
		newCollection = append(newCollection, operation(value))
	}
	return newCollection
}
