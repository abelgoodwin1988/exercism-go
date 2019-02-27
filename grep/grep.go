// Package grep contains function Search which accepts a patter nstring
//	flags, and a string of files. It will pattern match against lines
//	in provided files and return matches.
package grep

import (
	"bufio"
	"log"
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

// Use readable names for mappings
var printLineNumbers = flagsMap["-n"]
var printFileNames = flagsMap["-l"]
var invertMatch = flagsMap["-i"]
var ignoreCase = flagsMap["-v"]
var matchWholeLine = flagsMap["-x"]

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
	// set simple names for current search
	printLineNumbers = flagsMap["-n"]
	printFileNames = flagsMap["-l"]
	invertMatch = flagsMap["-i"]
	ignoreCase = flagsMap["-v"]
	matchWholeLine = flagsMap["-x"]
	// iterate through list of file names
	for _, file := range files {
		// Open the file
		f, err := os.Open(file)
		if err != nil {
			log.Fatalf("Error opening file %v: \n%v\n", file, err)
			return nil
		}
		defer f.Close()
		// Read the file
		scanner := bufio.NewScanner(f)
		var fileLineArray []string
		// Send each line to an array
		for scanner.Scan() {
			fileLineArray = append(fileLineArray, scanner.Text())
		}

		// Iterate through array, check for matches, assign to
		//	return string array.
		for lineNumber, value := range fileLineArray {
			if ignoreCase {
				if invertMatch && matchWholeLine {
					insensitivePattern := strings.ToLower(pattern)
					insensitiveValue := strings.ToLower(value)
					if insensitivePattern != insensitiveValue {
						matches = matched(matches, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if invertMatch {
					insensitivePattern := strings.ToLower(pattern)
					insensitiveValue := strings.ToLower(value)
					if !strings.Contains(insensitiveValue, insensitivePattern) {
						matches = matched(matches, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if matchWholeLine {
					if value != pattern {
						matches = matched(matches, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if !strings.Contains(value, pattern) {
					matches = matched(matches, file, value, lineNumber, multipleFiles)
				}
				continue
			} else {
				if invertMatch && matchWholeLine {
					insensitivePattern := strings.ToLower(pattern)
					insensitiveValue := strings.ToLower(value)
					if insensitivePattern == insensitiveValue {
						matches = matched(matches, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if invertMatch {
					insensitivePattern := strings.ToLower(pattern)
					insensitiveValue := strings.ToLower(value)
					if strings.Contains(insensitiveValue, insensitivePattern) {
						matches = matched(matches, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if matchWholeLine {
					if value == pattern {
						matches = matched(matches, file, value, lineNumber, multipleFiles)
					}
					continue
				}
				if strings.Contains(value, pattern) {
					matches = matched(matches, file, value, lineNumber, multipleFiles)
				}
				continue
			}
		}
	}
	return matches
}

func matched(matches []string, file string, value string, lineNumber int, multipleFiles bool) (rtnMatches []string) {
	var matchLine string
	if printFileNames {
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
