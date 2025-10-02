package student

func ConcatParams(args []string) string {
	result := ""
	for i, s := range args {
		if i > 0 {
			result += "\n"
		}
		result += s
	}
	return result
}
