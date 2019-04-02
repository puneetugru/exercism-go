package series

func All(n int, s string) (result []string) {
	for i := 0; n <= len(s); i++ {
		result = append(result, s[i:n])
		n++
	}
	return
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[:n], true
}
