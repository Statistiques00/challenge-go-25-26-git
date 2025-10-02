package student

func ToLower(s string) string {
	result := []rune{}
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			result = append(result, r-'A'+'a')
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}
