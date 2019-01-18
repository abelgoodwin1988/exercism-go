// Package raindrops implements a method to check factorials
//	of a number and if it contains 3, 5, or 7 convert to
//	a raindrop string and return, else return the number
//	 itsself
package raindrops

import "strconv"

var plingplangplog = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert converts numbers into raindrop sounds based on
//	the raindrops package doc block
func Convert(factor int) (raindrops string) {
	// series of ifs using modulus 0 to determine
	//	if the values are factorials
	for _, value := range []int{3, 5, 7} {
		if factor%value == 0 {
			raindrops += plingplangplog[value]
		}
	}
	if raindrops == "" {
		raindrops += strconv.Itoa(factor)
	}
	return raindrops
}
