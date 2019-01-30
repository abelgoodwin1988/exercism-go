// Package anagram containd method which returns a list of anagrams
package anagram

import (
	"strings"
	"unicode"
)

// Detect 's which words in an array of string are anagram
func Detect(match string, candidate []string) []string {
	success := []string{}
	matchMap := map[rune]int{}

	// declare the map
	setMapCount(&matchMap, &match)
	// iterate through candidates
	for _, value := range candidate {
		// check for same-character length
		if len(value) != len(match) {
			continue
		}
		// anagrams are not anagrams of themselves
		if strings.ToLower(value) == strings.ToLower(match) {
			continue
		}
		// Create rune map for candidate
		candidateMap := map[rune]int{}
		setMapCount(&candidateMap, &value)
		isAnagram := true
		for k, v := range matchMap {
			if _, ok := candidateMap[k]; ok {
				if v != candidateMap[k] {
					isAnagram = false
				}
			} else {
				isAnagram = false
			}
		}
		if isAnagram {
			success = append(success, value)
		}
	}
	return success
}

func setMapCount(pMap *map[rune]int, pMatch *string) {
	for _, val := range *pMatch {
		if _, ok := (*pMap)[unicode.ToLower(val)]; ok {
			(*pMap)[unicode.ToLower(val)]++
			continue
		}
		(*pMap)[unicode.ToLower(val)] = 1
	}
}
