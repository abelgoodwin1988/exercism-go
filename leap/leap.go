// Package leap exports a function which determines
//	whether or not a provided year is a leap year.
package leap

// IsLeapYear accepts an in representing a year of the
//	gregorian calendar and returns a a boolean for whether
//	it is or is not a leap year.
func IsLeapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false
}
