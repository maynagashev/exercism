package isbn

func IsValidISBN(isbn string) bool {
	var sum int
	var multiplier int = 10
	for _, c := range isbn {
		if c == 'X' && multiplier == 1 {
			sum += 10
			multiplier--
		} else if c >= '0' && c <= '9' {
			sum += int(c-'0') * multiplier
			multiplier--
		} else if c != '-' {
			// invalid character
			return false
		}
	}
	return sum%11 == 0 && multiplier == 0
}
