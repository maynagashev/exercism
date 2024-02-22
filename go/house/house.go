package house

// lines contains the unique parts of each verse in reverse order.
var lines = []string{
	"the house that Jack built.",
	"the malt\nthat lay in ",
	"the rat\nthat ate ",
	"the cat\nthat killed ",
	"the dog\nthat worried ",
	"the cow with the crumpled horn\nthat tossed ",
	"the maiden all forlorn\nthat milked ",
	"the man all tattered and torn\nthat kissed ",
	"the priest all shaven and shorn\nthat married ",
	"the rooster that crowed in the morn\nthat woke ",
	"the farmer sowing his corn\nthat kept ",
	"the horse and the hound and the horn\nthat belonged to ",
}

// Verse generates a single verse of the song.
func Verse(v int) string {
	verse := "This is "
	for i := v - 1; i >= 0; i-- {
		verse += lines[i]
	}
	return verse
}

// Song generates the entire song.
func Song() string {
	song := ""
	for i := 1; i <= len(lines); i++ {
		if i > 1 {
			song += "\n\n"
		}
		song += Verse(i)
	}
	return song
}
