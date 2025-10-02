package student

func ToUpper(s string) string {
	result := []rune{}
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			result = append(result, r-'a'+'A')
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
