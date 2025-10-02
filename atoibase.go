package student

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}
	b := len(base)
	res := 0
	for _, r := range s {
		idx := indexRune(base, r)
		if idx == -1 {
			return 0
		}
		res = res*b + idx
	}
	return res
}

func indexRune(s string, r rune) int {
	for i, v := range s {
		if v == r {
			return i
		}
	}
	return -1
}
