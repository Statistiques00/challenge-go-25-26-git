package student

func BasicAtoi2(s string) int {
	res := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		res = res*10 + int(r-'0')
	}
	return res
}
