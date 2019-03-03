// Package cryptosquare ...
package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

// Encode ...
func Encode(t string) string {
	// Normalize the input
	// Remove spaces, punctuation, & downcase the text
	t = strings.ToLower(t)
	normalizedText := []rune{}
	for _, r := range t {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			normalizedText = append(normalizedText, r)
		}
	}

	// Calcluate the number of columns and rows
	//	for the block cipher
	var col int
	var row int
	columnsRows := math.Sqrt(float64(len(normalizedText)))
	if columnsRows == float64(int(columnsRows)) {
		col, row = int(columnsRows), int(columnsRows)
	} else {
		row = int(math.Floor(columnsRows))
		col = row + 1
		if row*col < len(normalizedText) {
			row++
		}
	}

	// Construct a matrix where the normalized text overflows into
	//	the next row when the previous is filled.
	matrixText := make([][]rune, row)
	for i := 0; i < row; i++ {
		min := i * col
		max := (i + 1) * col
		if max > len(normalizedText)-1 {
			matrixText[i] = []rune(normalizedText[min:])
			continue
		}
		matrixText[i] = []rune(normalizedText[min:max])
	}
	// Iterate over the matrixText and construct the ciphered version
	//	by creating a [][]rune that contains the value read along the
	//	column
	cipheredRunes := make([][]rune, col)
	for i := 0; i < col; i++ {
		cipheredColText := make([]rune, row)
		for j := 0; j < row; j++ {
			if len(matrixText[j])-1 < i {
				cipheredColText[j] = ' '
				continue
			}
			cipheredColText[j] = matrixText[j][i]
		}
		cipheredRunes[i] = cipheredColText
	}
	// Flatten multidimensional array & return
	cipheredText := make([]rune, 0, row*col)
	for i, rs := range cipheredRunes {
		cipheredText = append(cipheredText, rs...)
		if i != len(cipheredRunes)-1 {
			cipheredText = append(cipheredText, ' ')
		}
	}
	return string(cipheredText)
}
