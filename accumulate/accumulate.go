package accumulate

const testVersion = 1

func Accumulate(str []string, f func(st string) string) (result []string) {
	for _, v := range str {
		result = append(result, []string{f(v)}...)
	}
	return result
}
