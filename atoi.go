package student

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	sign := 1
	i := 0
	if s[0] == '+' {
		sign = 1
		i++
	} else if s[0] == '-' {
		sign = -1
		i++
	}
	if i < len(s) && (s[i] == '+' || s[i] == '-') {
		return 0
	}
	res := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		res = res*10 + int(s[i]-'0')
	}
	return sign * res
}
