package secret

var enums = []string{"wink", "double blink", "close your eyes", "jump"}

// Handshake prints the code required for given number
func Handshake(code uint) (h []string) {
	switch {
	case code < 1:
	case code&16 == 0:
		for _, s := range enums {
			if code&1 != 0 {
				h = append(h, s)
			}
			code >>= 1
		}
	default:
		for i := 3; i >= 0; i-- {
			if code&8 != 0 {
				h = append(h, enums[i])
			}
			code <<= 1
		}
	}
	return
}
