package pangram

import "strings"

const testVersion = 2

func IsPangram(str string) bool {
	lower := strings.ToLower(str)
	var check [26]bool
	var count int

	for _, c := range lower {
		if c >= 'a' {
			if c > 'z' {
				continue
			}
			c -= 97
		} else {
			continue
		}

		if !check[c] {
			if count == 25 {
				return true
			}
			check[c] = true
			count++
		}
	}
	return false
}
