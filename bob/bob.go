// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

func yelling(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != strings.ToUpper(remark)
}

func asking(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func silent(remark string) bool {
	return remark == ""
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	switch remark = strings.TrimSpace(remark); {
	case silent(remark):
		return "Fine. Be that way!"
	case yelling(remark):
		if asking(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	case asking(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}
