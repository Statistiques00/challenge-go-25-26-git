package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	name := os.Args[0]
	last := -1
	for i, r := range name {
		if r == '/' || r == '\\' {
			last = i
		}
	}
	for _, r := range name[last+1:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
