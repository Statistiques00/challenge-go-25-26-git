package student

func Capitalize(s string) string {
	result := []rune{}
	inWord := false
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			if !inWord {
				if r >= 'a' && r <= 'z' {
					result = append(result, r-'a'+'A')
				} else {
					result = append(result, r)
				}
				inWord = true
			} else {
				if r >= 'A' && r <= 'Z' {
					result = append(result, r-'A'+'a')
				} else {
					result = append(result, r)
				}
			}
		} else {
			result = append(result, r)
			inWord = false
		}
	}
	return string(result)
}
