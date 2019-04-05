package letter

import (
	"sync"
)

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
//	of strings, merges the results of the frequencies and returns them.
func ConcurrentFrequency(sArr []string) FreqMap {
	m := FreqMap{}
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	for _, s := range sArr {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			for _, r := range s {
				mu.Lock()
				m[r]++
				mu.Unlock()
			}
		}(s)
	}
	wg.Wait()
	return m
}
