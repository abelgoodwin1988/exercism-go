// Package grains contains method Square and Total which
//	return the number of grains on an nth sqaure on a
//	chessboard, or a the sum of grains on a chessboard
package grains

import (
	"errors"
)

// Square accepts an integer representing the nth square
//	on a chessboard and returns a sum of powers representing
//	numbers of grains on the board following the pattern
//	in the readme.md
func Square(n int) (uint64, error) {
	// Catch out-of-bounds errors
	if n < 1 || n > 64 {
		return 0, errors.New("out-of-range chess-board square value")
	}
	return 1 << (uint(n) - 1), nil
}

// Total returns the total number of grains of rice
//	found on a chess-board following the doubling algorithm
func Total() uint64 {
	return uint64(1<<64 - 1)
}
