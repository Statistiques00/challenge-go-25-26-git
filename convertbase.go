package student

func ConvertBase(nbr, baseFrom, baseTo string) string {
	value := AtoiBase(nbr, baseFrom)
	if value == 0 {
		return string(baseTo[0])
	}
	b := len(baseTo)
	res := ""
	for value > 0 {
		res = string(baseTo[value%b]) + res
		value /= b
	}
	return res
}
