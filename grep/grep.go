// Package grep contains function Search which accepts a patter nstring
//	flags, and a string of files. It will pattern match against lines
//	in provided files and return matches.
package grep

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// flag readable names
var (
	printLineNumber      = "-n"
	printNameOnly        = "-l"
	matchCaseInsensitive = "-i"
	matchNonMatches      = "-v"
	matchEntireLine      = "-x"
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
			if strings.Contains(line, pattern) {
				match = append(match, line)
			}
		}

		if err := scanner.Err(); err != nil {
			log.Printf("Error reading file %v at line %v with error %v\n", fileName, i, err)
		}
	}
	return match
}
