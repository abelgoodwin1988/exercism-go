package pangram

import "strings"

const testVersion = 1

//IsPangram which determines if the sentence is Pangram
func IsPangram(sentence string) bool {

	var isPangram = true
	sentenceUpper := strings.ToUpper(sentence) //Converting the words to upper case (Pangram is case insensitive)

	for i := 65; i < 91; i++ { //ASCII Range of Upper Case
		if !strings.ContainsRune(sentenceUpper, rune(i)) {
			isPangram = false
			break
		}
	}
	return isPangram

}
