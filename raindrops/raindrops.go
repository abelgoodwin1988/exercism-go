// Package raindrops implements a method to check factorials
//	of a number and if it contains 3, 5, or 7 convert to
//	a raindrop string and return, else return the number
//	 itsself
package raindrops

import "strconv"

// Convert converts numbers into raindrop sounds based on
//	the raindrops package doc block
func Convert(factor int) (raindrops string) {
	// series of ifs using modulus 0 to determine
	//	if the values are factorials
	if factor%3 == 0 {
		raindrops += "Pling"
	}
	if factor%5 == 0 {
		raindrops += "Plang"
	}
	if factor%7 == 0 {
		raindrops += "Plong"
	}
	if raindrops == "" {
		raindrops += strconv.Itoa(factor)
	}
	return raindrops
}
