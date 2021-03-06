package house

import "strings"

var songLines = []string{
	"the horse and the hound and the horn\nthat belonged to",
	"the farmer sowing his corn\nthat kept",
	"the rooster that crowed in the morn\nthat woke",
	"the priest all shaven and shorn\nthat married",
	"the man all tattered and torn\nthat kissed",
	"the maiden all forlorn\nthat milked",
	"the cow with the crumpled horn\nthat tossed",
	"the dog\nthat worried",
	"the cat\nthat killed",
	"the rat\nthat ate",
	"the malt\nthat lay in",
}

//Verse function to create each line of the song
func Verse(v int) (verse string) {
	v--
	verse += "This is "
	verse += strings.Join(songLines[len(songLines)-v:], " ")
	if v > 0 {
		verse += " "
	}
	verse += "the house that Jack built."
	return
}

//Song function to create final song
func Song() (song string) {
	for i := 0; i <= len(songLines); i++ {
		song += Verse(i + 1)
		if i < len(songLines) {
			song += "\n\n"
		}
	}
	return
}
