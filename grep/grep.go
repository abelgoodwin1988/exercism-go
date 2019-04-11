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

// Search accepts a pattern, flags and filenames. Search will match the pattern
//	to the lines in the files while applying any matching rules found in the flags.
//	Flags:
// - `-n` Print the line numbers of each matching line.
// - `-l` Print only the names of files that contain at least one matching line.
// - `-i` Match line using a case-insensitive comparison.
// - `-v` Invert the program -- collect all lines that fail to match the pattern.
// - `-x` Only match entire lines, instead of lines that contain a match.
func Search(pattern string, flags []string, files []string) []string {
	// flag readable names
	var (
		displayLineNumber     = false
		displayFileNameOnly   = false
		matchCaseInsensitive  = false
		matchNonMatches       = false
		matchEntireLine       = false
		multipleFilesMatching = false
	)
	// Parse flags into readable names
	for _, flag := range flags {
		switch flag {
		case "-n":
			displayLineNumber = true
		case "-l":
			displayFileNameOnly = true
		case "-i":
			matchCaseInsensitive = true
		case "-v":
			matchNonMatches = true
		case "-x":
			matchEntireLine = true
		}
	}
	if len(files) > 1 {
		multipleFilesMatching = true
	}

	// Begin matching
	match := []string{}
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			log.Printf("Error opening file %v with error %v\n", fileName, err)
		}
		defer file.Close()

		// Read each line by line and evaluate for matches based on flags
		scanner := bufio.NewScanner(file)
		i := 0
		for scanner.Scan() {
			line := scanner.Text()
			matchesAtStartOfIteration := len(match)
			// default case followed serialized cases
			//
			if !matchCaseInsensitive && !matchNonMatches && !matchEntireLine {
				if strings.Contains(line, pattern) {
					match = append(match, line)
				}
			}
			// insensitive match
			if matchCaseInsensitive && !matchNonMatches && !matchEntireLine {
				if strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
					match = append(match, line)
				}
			}
			// insensitive match & nonMatch
			if matchCaseInsensitive && matchNonMatches && !matchEntireLine {
				if !strings.Contains(strings.ToLower(line), strings.ToLower(pattern)) {
					match = append(match, line)
				}
			}
			// insensitive match             & entire line match
			if matchCaseInsensitive && !matchNonMatches && matchEntireLine {
				if strings.ToLower(line) == strings.ToLower(pattern) {
					match = append(match, line)
				}
			}
			// insensitive match & nonMatch & entire line match
			if matchCaseInsensitive && matchNonMatches && matchEntireLine {
				if strings.ToLower(line) != strings.ToLower(pattern) {
					match = append(match, line)
				}
			}
			//                     nonMatch
			if !matchCaseInsensitive && matchNonMatches && !matchEntireLine {
				if !strings.Contains(line, pattern) {
					match = append(match, line)
				}
			}
			//                     nonMatch & entire line match
			if !matchCaseInsensitive && matchNonMatches && matchEntireLine {
				if line != pattern {
					match = append(match, line)
				}
			}
			//                                entire line match
			if !matchCaseInsensitive && !matchNonMatches && matchEntireLine {
				if line == pattern {
					match = append(match, line)
				}
			}

			matchesAtEndOfIteration := len(match)
			lineWasMatched := matchesAtStartOfIteration != matchesAtEndOfIteration
			// "Print" additions and rules
			// add line display number
			if displayLineNumber && lineWasMatched {
				match[len(match)-1] = strconv.Itoa(i+1) + ":" + match[len(match)-1]
			}
			// add filename to start if we're iterating over multiple files
			if multipleFilesMatching && lineWasMatched {
				match[len(match)-1] = fileName + ":" + match[len(match)-1]
			}
			if displayFileNameOnly && lineWasMatched {
				match[len(match)-1] = fileName
				// dedupe
				encountered := map[string]struct{}{}
				for _, encounter := range match {
					encountered[encounter] = struct{}{}
				}
				match = []string{}
				for encounter := range encountered {
					match = append(match, encounter)
				}
			}
			// increment line tracker
			i++
		}

		if err := scanner.Err(); err != nil {
			log.Printf("Error reading file %v at line %v with error %v\n", fileName, i, err)
		}
	}
	return match
}
