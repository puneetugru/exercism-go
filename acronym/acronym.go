package acronym

import (
	"regexp"
	"fmt"
	"strings"
)

const testVersion = 3

func Abbreviate(str string) string {
	regex 	:= regexp.MustCompile("[A-Z]+[a-z]*|[a-z]+")
	words 	:= regex.FindAllString(str, -1)
	abbr 	:= []string{}

	for _, word := range words {
		abbr = append(abbr, string(word[0]))
	}

	return fmt.Sprintf("%s", strings.ToUpper(strings.Join(abbr, "")))
}
