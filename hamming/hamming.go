package hamming

import "errors"

const testVersion = 6

func Distance(a, b string) (ret int, err error) {
	if len(b) != len(a) {
		return 0, errors.New("different length strings")
	}

	for i:=0; i<len(a); i++ {
		if a[i] != b[i] {
			ret++
		}
	}
	return
}
