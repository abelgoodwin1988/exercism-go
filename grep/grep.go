// Package grep contains function Search which accepts a patter nstring
//	flags, and a string of files. It will pattern match against lines
//	in provided files and return matches.
package grep

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// create constant holding map of false strings; when true we'll
//	enable the features for the flags
var flagsMap = map[string]bool{
	"-n": false,
	"-l": false,
	"-i": false,
	"-v": false,
	"-x": false,
}

/*
Search accepts a pattern string, flags string and  files []string.

Search will iterate through the files in the order provided and add any
matching lines of text to the return.

If the search had more than one file, each string should have the
filename as a suffix to the matching line.

-- Flags
- -n Print the line numbers of each matching line.
- -l Print only the names of files that contain at least one matching line.
- -v Invert the program -- collect all lines that fail to match the pattern.
- -i Match line using a case-insensitive comparison.
- -x Only match entire lines, instead of lines that contain a match.
*/
func Search(pattern string, flags []string, files []string) (matches []string) {
	// instantiate empty matches for match
	matches = []string{}
	// multiple files?
	multipleFiles := len(files) > 1
	// reset mapping
	for i := range flagsMap {
		flagsMap[i] = false
	}
	// set mapping for current search
	for _, value := range flags {
		flagsMap[value] = true
	}
	// Initalize the return array of string

	// iterate through list of file names
	for _, file := range files {
		// Open the file
		f, err := os.Open(file)
		defer f.Close()
		check(err)
		// Read the file
		scanner := bufio.NewScanner(f)
		var fileArray []string
		// Send each line to an array
		for scanner.Scan() {
			fileArray = append(fileArray, scanner.Text())
		}

		// Iterate through array, check for matches, assign to
		//	return string array.
		for lineNumber, value := range fileArray {
			if flagsMap["-v"] {
				if flagsMap["-i"] && flagsMap["-x"] {
					insensitivePattern := strings.ToLower(pattern)
					insensitiveValue := strings.ToLower(value)
					if insensitivePattern != insensitiveValue {
						matches = matched(matches, flagsMap, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if flagsMap["-i"] {
					insensitivePattern := strings.ToLower(pattern)
					insensitiveValue := strings.ToLower(value)
					if !strings.Contains(insensitiveValue, insensitivePattern) {
						matches = matched(matches, flagsMap, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if flagsMap["-x"] {
					if value != pattern {
						matches = matched(matches, flagsMap, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if !strings.Contains(value, pattern) {
					matches = matched(matches, flagsMap, file, value, lineNumber, multipleFiles)
				}
				continue
			} else {
				if flagsMap["-i"] && flagsMap["-x"] {
					insensitivePattern := strings.ToLower(pattern)
					insensitiveValue := strings.ToLower(value)
					if insensitivePattern == insensitiveValue {
						matches = matched(matches, flagsMap, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if flagsMap["-i"] {
					insensitivePattern := strings.ToLower(pattern)
					insensitiveValue := strings.ToLower(value)
					if strings.Contains(insensitiveValue, insensitivePattern) {
						matches = matched(matches, flagsMap, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if flagsMap["-x"] {
					if value == pattern {
						matches = matched(matches, flagsMap, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if strings.Contains(value, pattern) {
					matches = matched(matches, flagsMap, file, value, lineNumber, multipleFiles)
				}
				continue
			}
		}
	}
	return matches
}

func matched(matches []string, flagsMap map[string]bool, file string, value string, lineNumber int, multipleFiles bool) (rtnMatches []string) {
	var matchLine string
	if flagsMap["-l"] {
		for _, fileName := range matches {
			if fileName == file {
				return matches
			}
		}
		matchLine += file
		rtnMatches = append(matches, matchLine)
		return rtnMatches
	}
	if multipleFiles {
		matchLine += file + ":"

	}
	if flagsMap["-n"] {
		matchLine += strconv.Itoa(lineNumber+1) + ":"
	}
	if multipleFiles {
		matchLine += value
		rtnMatches = append(matches, matchLine)
		return rtnMatches
	}

	matchLine += value
	rtnMatches = append(matches, matchLine)
	return rtnMatches
}

// simple err check w/ panic.
func check(e error) {
	if e != nil {
		panic(e)
	}
}
