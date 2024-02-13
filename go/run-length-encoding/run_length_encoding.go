package encode

import (
	"strconv"
	"strings"
)

func RunLengthEncode(input string) string {
	var result strings.Builder
	count := 1

	for i := 0; i < len(input); i++ {
		if i+1 < len(input) && input[i] == input[i+1] {
			count++
		} else {
			if count > 1 {
				result.WriteString(strconv.Itoa(count))
			}
			result.WriteByte(input[i])
			count = 1
		}
	}

	return result.String()
}

func RunLengthDecode(input string) string {
	var result strings.Builder
	count := 0

	for i := 0; i < len(input); i++ {
		if input[i] >= '0' && input[i] <= '9' {
			count = count*10 + int(input[i]-'0')
		} else {
			if count == 0 {
				count = 1
			}
			result.WriteString(strings.Repeat(string(input[i]), count))
			count = 0
		}
	}

	return result.String()
}
