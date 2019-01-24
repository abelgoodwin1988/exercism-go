// Package diffsquares takes a
package diffsquares

import "math"

// SquareOfSum takes an int and return the square of the sum
func SquareOfSum(n int) int {
	var squareOfSums int
	for i := 1; i <= n; i++ {
		squareOfSums += i
	}
	return int(math.Pow(float64(squareOfSums), 2))
}

// SumOfSquares takes an int and returns the sum of n's squared
func SumOfSquares(n int) int {
	var sumOfSquares int
	for i := 1; i <= n; i++ {
		sumOfSquares += int(math.Pow(float64(i), 2))
	}
	return sumOfSquares
}

// Difference returns the difference between the two calculations
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
