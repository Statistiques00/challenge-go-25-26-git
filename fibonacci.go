package student

func Fibonacci(n int) int {
	if n < 0 {
		return -1
	} else if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
