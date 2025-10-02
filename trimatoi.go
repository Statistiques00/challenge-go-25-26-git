package student

func TrimAtoi(s string) int {
	sign := 1
	started := false
	res := 0
	for _, r := range s {
		if !started && r == '-' {
			sign = -1
		}
		if r >= '0' && r <= '9' {
			res = res*10 + int(r-'0')
			started = true
		}
	}
	return sign * res
}
