// Package pythagorean contains method(s) Sum which sums a
//	pythagorean triplet
package pythagorean

// Triplet is a struct representing the value used
//	in a pythagorean triplet
type Triplet struct {
	a int
	b int
	c int
}

// Sum accepts a sum of pythagorean a triplet. It then calculates
//	and returns all possible pythagorean triplets.
func Sum(p int) []Triplet {
	t := []Triplet{}
	for a := 2; a < p/3; a++ {
		for b := a + 1; b < p/2; b++ {
			c := p - a - b
			if a*a+b*b == c*c {
				t = append(t, Triplet{a, b, c})
			}
		}
	}
	return t
}

// Range ...
func Range(min, max int) []Triplet {
	t := []Triplet{}
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			for c := b + 1; c <= max; c++ {
				if a*a+b*b == c*c {
					t = append(t, Triplet{a, b, c})
				}
			}
		}
	}
	return t
}
