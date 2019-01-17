// Package triangle accepts three lengths and returns
//	what type of triangle the three lengths would
//	make
package triangle

import "math"

// Kind implements string const values
type Kind string

// Declaration of const values to be represented
//	in testing
const (
	NaT = "Nat"
	Equ = "Equ"
	Iso = "Iso"
	Sca = "Sca"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	// Check that each side has a non-zero, unsigned value
	if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
		return k
	}
	// Check that we have -at least- a degenerate triangle
	if a+b < c ||
		a+c < b ||
		b+c < a {
		k = NaT
		return k
	}
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) {
		k = NaT
		return k
	}
	if math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		k = NaT
		return k
	}
	if a == b && b == c {
		k = Equ
	} else if (a == b && a != c) ||
		(a == c && a != b) ||
		(b == c && b != a) {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
