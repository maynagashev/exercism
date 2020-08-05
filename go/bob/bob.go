/*
Package bob is about a lackadaisical teenager. In conversation, his responses are very limited.

Bob answers 'Sure.' if you ask him a question, such as "How are you?".
He answers 'Whoa, chill out!' if you YELL AT HIM (in all capitals).
He answers 'Calm down, I know what I'm doing!' if you yell a question at him.
He says 'Fine. Be that way!' if you address him without actually saying
anything.

He answers 'Whatever.' to anything else.
*/
package bob

import (
	"strings"
	"unicode"
)

// Hey builds answer to phrase
func Hey(remark string) (res string) {
	switch {
	case isEmpty(remark):
		res = "Fine. Be that way!"
	case isAllCapitals(remark) && isQuestion(remark):
		res = "Calm down, I know what I'm doing!"
	case isQuestion(remark):
		res = "Sure."
	case isAllCapitals(remark):
		res = "Whoa, chill out!"
	default:
		res = "Whatever."
	}
	return
}

func isEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

func isAllCapitals(s string) bool {
	hasLetters := strings.IndexFunc(s, unicode.IsLetter) >= 0
	return hasLetters && strings.ToUpper(s) == s
}

func isQuestion(s string) bool {
	return strings.HasSuffix(strings.TrimSpace(s), "?")
}
