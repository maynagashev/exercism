package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Search searches for a pattern in the given files according to the specified flags.
func Search(pattern string, flags, files []string) []string {
	var results []string
	caseInsensitive := contains(flags, "-i")
	invertMatch := contains(flags, "-v")
	matchEntireLine := contains(flags, "-x")
	printLineNumbers := contains(flags, "-n")
	printFileNames := contains(flags, "-l")

	for _, file := range files {
		fileMatches := false
		lineNumber := 1
		f, err := os.Open(file)
		if err != nil {
			fmt.Printf("Error opening file: %s\n", err)
			continue
		}
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			match := matches(line, pattern, caseInsensitive, matchEntireLine)
			if invertMatch {
				match = !match
			}
			if match {
				fileMatches = true
				if printFileNames {
					results = append(results, file)
					break
				} else {
					prefix := ""
					if len(files) > 1 {
						prefix = file + ":"
					}
					if printLineNumbers {
						prefix += fmt.Sprintf("%d:", lineNumber)
					}
					results = append(results, prefix+line)
				}
			}
			lineNumber++
		}
		f.Close()
		if printFileNames && fileMatches && !contains(results, file) {
			results = append(results, file)
		}
	}
	return results
}

// contains checks if a slice contains a string.
func contains(slice []string, str string) bool {
	for _, item := range slice {
		if item == str {
			return true
		}
	}
	return false
}

// matches checks if a line matches the pattern according to flags.
func matches(line, pattern string, caseInsensitive, matchEntireLine bool) bool {
	if caseInsensitive {
		line = strings.ToLower(line)
		pattern = strings.ToLower(pattern)
	}
	if matchEntireLine {
		return line == pattern
	}
	return strings.Contains(line, pattern)
}
