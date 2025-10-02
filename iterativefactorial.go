package student

func IterativeFactorial(n int) int {
	if n < 0 {
		return 0
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		if result < 0 {
			return 0
		}
	}
	return result
}
