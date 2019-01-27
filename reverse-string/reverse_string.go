// Package reverse ...
package reverse

// String accepts a string and reverses it
func String(s string) string {
	var r string
	for _, value := range s {
		r = string(value) + r
	}
	return r
}
