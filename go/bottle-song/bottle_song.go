package bottlesong

import (
	"fmt"
	"strings"
)

/*
Ten green bottles hanging on the wall,
Ten green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be nine green bottles hanging on the wall.

Nine green bottles hanging on the wall,
Nine green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be eight green bottles hanging on the wall.

Eight green bottles hanging on the wall,
Eight green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be seven green bottles hanging on the wall.

Seven green bottles hanging on the wall,
Seven green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be six green bottles hanging on the wall.

Six green bottles hanging on the wall,
Six green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be five green bottles hanging on the wall.

Five green bottles hanging on the wall,
Five green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be four green bottles hanging on the wall.

Four green bottles hanging on the wall,
Four green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be three green bottles hanging on the wall.

Three green bottles hanging on the wall,
Three green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be two green bottles hanging on the wall.

Two green bottles hanging on the wall,
Two green bottles hanging on the wall,
And if one green bottle should accidentally fall,
There'll be one green bottle hanging on the wall.

One green bottle hanging on the wall,
One green bottle hanging on the wall,
And if one green bottle should accidentally fall,
There'll be no green bottles hanging on the wall.
*/

// Recite the lyrics to that popular children's repetitive song: Ten Green Bottles.
func Recite(startBottles, takeDown int) []string {
	var song []string
	for i := 0; i < takeDown; i++ {
		if startBottles == 0 {
			break
		}
		song = append(song, fmt.Sprintf("%s hanging on the wall,", UpperCaseFirstLetter(bottle(startBottles))))
		song = append(song, fmt.Sprintf("%s hanging on the wall,", UpperCaseFirstLetter(bottle(startBottles))))
		song = append(song, "And if one green bottle should accidentally fall,")
		startBottles--
		if startBottles == 0 {
			song = append(song, "There'll be no green bottles hanging on the wall.")
		} else {
			song = append(song, fmt.Sprintf("There'll be %s hanging on the wall.", bottle(startBottles)))
		}

		// Add an empty line between verses
		if i < takeDown-1 {
			song = append(song, "")
		}
	}
	return song
}

func bottle(bottles int) string {
	switch bottles {
	case 0:
		return "no green bottles"
	case 1:
		return "one green bottle"
	case 2:
		return "two green bottles"
	case 3:
		return "three green bottles"
	case 4:
		return "four green bottles"
	case 5:
		return "five green bottles"
	case 6:
		return "six green bottles"
	case 7:
		return "seven green bottles"
	case 8:
		return "eight green bottles"
	case 9:
		return "nine green bottles"
	case 10:
		return "ten green bottles"
	default:
		return fmt.Sprintf("%d bottles", bottles)
	}
}

func UpperCaseFirstLetter(s string) string {
	words := strings.Fields(s)
	if len(words) > 0 {
		words[0] = strings.Title(words[0])
	}
	return strings.Join(words, " ")
}
