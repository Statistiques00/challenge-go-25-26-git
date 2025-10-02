package main

import (
	"os"
)

func atoi(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}
	sign := int64(1)
	start := 0
	if s[0] == '-' {
		sign = -1
		start = 1
	}
	var n int64
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		d := int64(s[i] - '0')
		if n > (1<<63-1-d)/10 {
			return 0, false
		}
		n = n*10 + d
	}
	return sign * n, true
}

func printInt(n int64) {
	if n == 0 {
		os.Stdout.Write([]byte{'0'})
		os.Stdout.Write([]byte{'\n'})
		return
	}
	buf := [21]byte{}
	neg := n < 0
	if neg {
		n = -n
	}
	i := len(buf)
	for n > 0 {
		i--
		buf[i] = byte(n%10) + '0'
		n /= 10
	}
	if neg {
		i--
		buf[i] = '-'
	}
	os.Stdout.Write(buf[i:])
	os.Stdout.Write([]byte{'\n'})
}

func printStr(s string) {
	os.Stdout.Write([]byte(s))
	os.Stdout.Write([]byte{'\n'})
}

func main() {
	if len(os.Args) != 4 {
		return
	}
	aStr, op, bStr := os.Args[1], os.Args[2], os.Args[3]
	a, okA := atoi(aStr)
	b, okB := atoi(bStr)
	if !okA || !okB {
		return
	}
	switch op {
	case "+":
		res := a + b
		if (b > 0 && res < a) || (b < 0 && res > a) {
			return
		}
		printInt(res)
	case "-":
		res := a - b
		if (b < 0 && res < a) || (b > 0 && res > a) {
			return
		}
		printInt(res)
	case "*":
		res := a * b
		if b != 0 && res/b != a {
			return
		}
		printInt(res)
	case "/":
		if b == 0 {
			printStr("No division by 0")
			return
		}
		printInt(a / b)
	case "%":
		if b == 0 {
			printStr("No modulo by 0")
			return
		}
		printInt(a % b)
	default:
		return
	}
}
