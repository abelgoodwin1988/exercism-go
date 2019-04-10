package letter

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency runs concurrent frequency against an array
//	of strings, merges the results of the FreqMap, and returns merged result.
func ConcurrentFrequency(sArr []string) FreqMap {
	m := FreqMap{}
	fmChan := make(chan FreqMap)

	for _, s := range sArr {
		go func(s string) {
			fmChan <- Frequency(s)
		}(s)
	}
	var fm FreqMap
	for i := 0; i < len(sArr); i++ {
		fm = <-fmChan
		for key, val := range fm {
			m[key] = m[key] + val
		}
	}
	close(fmChan)
	return m
}
