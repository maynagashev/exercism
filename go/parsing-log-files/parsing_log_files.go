// Package parsinglogfiles provides functions for parsing and manipulating log files.
package parsinglogfiles

import (
	"regexp"
)

// IsValidLine checks if a given line of text is valid based on a specific pattern.
// It returns true if the line starts with one of the following log levels: TRC, DBG, INF, WRN, ERR, FTL.
// Otherwise, it returns false.
func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`)
	return re.MatchString(text)
}

// SplitLogLine splits a given line of text into a slice of strings.
// It uses a regular expression to split the line at each occurrence of a pattern.
// The pattern is a sequence of characters enclosed in angle brackets.
func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	return re.Split(text, -1)
}

// CountQuotedPasswords counts the number of quoted passwords in a slice of strings.
// It uses a case-insensitive regular expression to find all occurrences of the word "password" enclosed in quotes.
// It returns the total count of such occurrences.
func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	count := 0
	for _, line := range lines {
		m := re.FindAllString(line, -1)
		if m != nil {
			count += len(m)
		}
	}
	return count
}

// RemoveEndOfLineText removes the end-of-line text from a given line of text.
// It uses a regular expression to find and replace all occurrences of the pattern "end-of-line" followed by one or more digits.
// It returns the modified line of text.
func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
	return re.ReplaceAllString(text, "")
}

// TagWithUserName tags each line of text in a slice of strings with the username.
// It uses a regular expression to find the username in each line.
// If a username is found, it prepends the line with a tag containing the username.
// It returns the modified slice of strings.
func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\S+)`)
	for i, line := range lines {
		if re.MatchString(line) {
			m := re.FindStringSubmatch(line)
			lines[i] = "[USR] " + m[1] + " " + line
		}
	}
	return lines
}
