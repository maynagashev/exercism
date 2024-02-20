package foodchain

import "strings"

// animals and reactions contain the unique parts of each verse in order.
var animals = []string{
	"fly",
	"spider",
	"bird",
	"cat",
	"dog",
	"goat",
	"cow",
	"horse",
}

var reactions = []string{
	"",
	"It wriggled and jiggled and tickled inside her.",
	"How absurd to swallow a bird!",
	"Imagine that, to swallow a cat!",
	"What a hog, to swallow a dog!",
	"Just opened her throat and swallowed a goat!",
	"I don't know how she swallowed a cow!",
	"She's dead, of course!",
}

// Verse generates a verse of the song.
func Verse(v int) string {
	if v == 8 {
		return "I know an old lady who swallowed a horse.\nShe's dead, of course!"
	}

	verse := []string{"I know an old lady who swallowed a " + animals[v-1] + "."}
	if reactions[v-1] != "" {
		verse = append(verse, reactions[v-1])
	}
	for i := v - 1; i > 0; i-- {
		line := "She swallowed the " + animals[i] + " to catch the " + animals[i-1]
		if animals[i-1] == "spider" {
			line += " that wriggled and jiggled and tickled inside her"
		}
		verse = append(verse, line+".")
	}
	verse = append(verse, "I don't know why she swallowed the fly. Perhaps she'll die.")

	return strings.Join(verse, "\n")
}

// Verses generates a range of verses of the song.
func Verses(start, end int) string {
	var verses []string
	for i := start; i <= end; i++ {
		verses = append(verses, Verse(i))
		if i < end {
			verses = append(verses, "\n\n")
		}
	}
	return strings.Join(verses, "")
}

// Song generates the entire song.
func Song() string {
	return Verses(1, 8)
}
