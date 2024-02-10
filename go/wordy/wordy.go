/*
Package wordy parses and evaluates simple math word problems returning the answer as an integer.
Handle a set of operations, in sequence.


Since these are verbal word problems, evaluate the expression from left-to-right, ignoring the typical order of operations.

What is 5 plus 13 plus 6?
24

What is 3 plus 2 multiplied by 3?
15 (i.e. not 9)

Problems with no operations simply evaluate to the number given.

What is 5?
Evaluates to 5.
*/

package wordy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Answer(question string) (res int, ok bool) {

	// remove What is from the question & remove the question mark
	re := regexp.MustCompile(`^What is (.*)\?$`)
	m := re.FindStringSubmatch(question)
	if len(m) != 2 {
		return 0, false
	}
	question = m[1]

	// find the first number in the question
	re = regexp.MustCompile(`^-?\d+`)
	asciiNum := re.FindString(question)
	if asciiNum != "" {
		num, err := strconv.Atoi(asciiNum)
		if err != nil {
			fmt.Printf("error converting number: '%s', error: %s\n", m, err)
			return 0, false
		}
		res += num
	}
	question = strings.TrimSpace(re.ReplaceAllString(question, ""))

	// find the operations in the question
	re = regexp.MustCompile(`^(plus|minus|multiplied by|divided by)\s*(-?\d+)`)
	for re.FindString(question) != "" {
		m := re.FindStringSubmatch(question)

		fmt.Printf("question: %s, matches: %+v\n", question, m)
		if len(m) != 3 {
			return 0, false
		}
		operation := m[1]
		num, err := strconv.Atoi(m[2])
		if err != nil {
			return 0, false
		}
		switch operation {
		case "plus":
			res += num
		case "minus":
			res -= num
		case "multiplied by":
			res *= num
		case "divided by":
			res /= num
		}
		// remove the operation from the question
		question = re.ReplaceAllString(question, "")
		question = strings.TrimSpace(question)
	}

	// if there is anything left in the question, it's an error
	if question != "" {
		return 0, false
	}

	return res, true
}
