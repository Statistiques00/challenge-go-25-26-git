package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}
	if len(args) == 0 {
		return
	}
	for _, s := range args {
		n := atoi(s)
		if n >= 1 && n <= 26 {
			var r rune
			if upper {
				r = rune('A' + n - 1)
			} else {
				r = rune('a' + n - 1)
			}
			z01.PrintRune(r)
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}

func atoi(s string) int {
	res := 0
	if s == "" {
		return 0
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		res = res*10 + int(r-'0')
	}
	return res
}
