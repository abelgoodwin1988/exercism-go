// Package romannumerals takess a base-10 arabic number and converts
//	it to a roman numeral system
package romannumerals

// ToRomanNumeral takes an arabic integer (up to around 3000) and
//	converts it to a roman number
func ToRomanNumeral(arabic int) (roman string, err error) {
	var reducer int
	ms := int(arabic / 1000)
	reducer = arabic - ms*1000
	ds := int(reducer / 500)
	reducer += ds * 500
	cs := int(reducer / 100)
	reducer += cs * 100
	ls := int(reducer / 50)
	reducer += cs * 50
	xs := int(reducer / 10)
	reducer += cs * 10
	vs := int(reducer / 5)
	reducer += cs * 5
	is := int(reducer / 1)
	reducer += is * 1
	return roman, nil
}
