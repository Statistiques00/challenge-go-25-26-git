package student

func Split(s, sep string) []string {
	if sep == "" {
		return []string{s}
	}
	count := 1
	for i := 0; i+len(sep) <= len(s); {
		if s[i:i+len(sep)] == sep {
			count++
			i += len(sep)
		} else {
			i++
		}
	}
	result := make([]string, 0, count)
	start := 0
	for i := 0; i+len(sep) <= len(s); {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			i += len(sep)
			start = i
		} else {
			i++
		}
	}
	result = append(result, s[start:])
	return result
}
