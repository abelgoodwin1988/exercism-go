// Package provide includes a functin Proverb which
//	takes an arbitrary numbers of string
//	and constructs a proverb.
package proverb

// Proverb takes an arbitrary number of strings and constructs
//	a proverb.
func Proverb(rhyme []string) []string {
	proverb := []string{}
	// catch zero-length case
	if len(rhyme) == 0 {
		return proverb
	}
	for i, value := range rhyme {
		if i < len(rhyme)-1 {
			proverb = append(proverb, "For want of a "+value+" the "+rhyme[i+1]+" was lost.")
		} else {
			proverb = append(proverb, "And all for the want of a "+rhyme[0]+".")
		}
	}
	return proverb
}
