package pascal

func Triangle(n int) (out [][]int) {
	if n < 1 {
		return
	}

	out = make([][]int, n)
	r := []int{1}

	out[0] = r
	for i := 1; i < n; i++ {
		last := r
		r = make([]int, i+1)
		r[0] = 1
		r[i] = 1
		for j := 1; j < i; j++ {
			r[j] = last[j-1] + last[j]
		}
		out[i] = r
	}

	return out
}
