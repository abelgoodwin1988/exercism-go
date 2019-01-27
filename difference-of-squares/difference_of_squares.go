// Package diffsquares takes a
package diffsquares

import "math"

// SquareOfSum takes an int and return the square of the sum
func SquareOfSum(n int) int {
	return int(math.Pow(float64(n*(n+1)/2), 2.0))
}

// SumOfSquares takes an int and returns the sum of n's squared
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference returns the difference between the two calculations
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
