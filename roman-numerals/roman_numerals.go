// Package romannumerals takess a base-10 arabic number and converts
//	it to a roman numeral system
package romannumerals

import "errors"

// numerals is used to track how many 10-based numerals
//	are present in an arabic number
var numerals = map[rune]int{
	'M': 0,
	'C': 0,
	'X': 0,
	'I': 0,
}

// numeralsOrder is a descending order of available roman numerals
var numeralsOrder = []rune{'M', 'D', 'C', 'L', 'X', 'V', 'I'}

// ToRomanNumeral takes an arabic integer (up to around 3000) and
//	converts it to a roman number
func ToRomanNumeral(arabic int) (roman string, err error) {
	var reducer int
	// capture any errors, or out-of-scope arabic numers
	if arabic < 1 || arabic > 3000 {
		return "", errors.New("out-of-range arabic")
	}
	// load the numerals map with the values belonging to each
	//	'main' 10-based numeral value
	numerals['M'] = int(arabic / 1000)
	reducer = arabic - numerals['M']*1000
	numerals['C'] = int(reducer / 100)
	reducer -= numerals['C'] * 100
	numerals['X'] = int(reducer / 10)
	reducer -= numerals['X'] * 10
	numerals['I'] = int(reducer / 1)
	reducer -= numerals['I'] * 1
	// Construct the roman number in another function
	roman = constructRoman()
	return roman, nil
}

// constructRoman uses globally defined 'numerals' mapping of
//	a count of 10-based roman numerals count in an arabic
//	number, gets the roman numeral arrangement and stores
//	in the respective positions.
func constructRoman() (roman string) {
	var romanPosition [4]string
	romanPosition[0] = getXNumerals('M', numerals['M'])
	romanPosition[1] = getNumeralBuild(numerals['C'], 2)
	romanPosition[2] = getNumeralBuild(numerals['X'], 4)
	romanPosition[3] = getNumeralBuild(numerals['I'], 6)
	for _, value := range romanPosition {
		roman += value
	}
	return roman
}

// getNumeralBuild takes an arbitrary count of a certain roman
//	numeral and a position of that roman numeral within a
//	descending array of roman numeral values
// It then returns the build for that single position/scope of
//	roman numeral
func getNumeralBuild(count int, position int) (numerals string) {
	if count == 9 {
		return string(numeralsOrder[position]) + string(numeralsOrder[position-2])
	}
	if count >= 5 {
		halfs := int(count / 5)
		remainders := count % 5
		numerals = getXNumerals(numeralsOrder[position-1], halfs)
		numerals += getXNumerals(numeralsOrder[position], remainders)
		return numerals
	}
	if count == 4 {
		return string(numeralsOrder[position]) + string(numeralsOrder[position-1])
	}
	numerals = getXNumerals(numeralsOrder[position], count)
	return numerals
}

// getXNumerals takes a rune and returns a string with that rune
//	x times reapeated.
func getXNumerals(numeral rune, count int) (numerals string) {
	for i := 0; i < count; i++ {
		numerals += string(numeral)
	}
	return numerals
}
