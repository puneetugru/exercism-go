package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}
	for _, word := range strings.Fields(normalize(phrase)) {
		word = strings.Trim(word, "'")
		freq[word]++
	}
	return freq
}

func normalize(phrase string) string {
	//  Allow for apostrophes in words.
	r, _ := regexp.Compile(`[^\w|']`)
	phrase = strings.ToLower(phrase)
	return r.ReplaceAllLiteralString(phrase, " ")
}
