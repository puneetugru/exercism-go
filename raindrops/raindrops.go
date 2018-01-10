package raindrops

import "strconv"

const testVersion = 3

func Convert(num int) string {
	str := ""
	if (num%3 ==0) {
		str = str+"Pling"
	}
	if (num%5 ==0) {
		str = str+"Plang"
	}
	if (num%7 ==0) {
		str = str+"Plong"
	}
	if len(str) == 0 {
		str = strconv.Itoa(num)
	}
	return str
}

// Don't forget the test program has a benchmark too.
// How fast does your Convert convert?
