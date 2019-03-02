// Package summultiples ...
package summultiples

// SumMultiples accepts a limit value and divisors of that value.
//	The divisors represent multiples of itsself up to a limit of
//	the limit variable. Taking the unique multiples of the
//	divisors, return their sum.
func SumMultiples(limit int, divisors ...int) int {
	// Map of integer with empty struct; empty struct so-as to not
	//	use extra memory for the keys value.
	multiples := make(map[int]struct{})
	var sum int // hold sum of multiples
	// Iterate through divisors, get count of multiples up to the limit
	//	and store the multiples to the multiples map
	for _, divisor := range divisors {
		// catch division by 0
		if divisor == 0 {
			continue
		}
		multiplesInLimit := limit / divisor
		for index := 1; index <= multiplesInLimit; index++ {
			if index*divisor >= limit {
				continue
			}
			// Only sum non-duplicative multiples from all multiples
			//	of divisor multiples
			if _, ok := multiples[index*divisor]; !ok {
				multiples[index*divisor] = struct{}{}
				sum += index * divisor
			}
		}
	}
	return sum
}
