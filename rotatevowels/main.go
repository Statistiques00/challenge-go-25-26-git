package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'A' ||
		r == 'e' || r == 'E' ||
		r == 'i' || r == 'I' ||
		r == 'o' || r == 'O' ||
		r == 'u' || r == 'U'
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}
	vowels := []rune{}
	for _, arg := range args {
		for _, r := range arg {
			if isVowel(r) {
				vowels = append(vowels, r)
			}
		}
	}
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}
	vi := 0
	for ai, arg := range args {
		for _, r := range arg {
			if isVowel(r) && len(vowels) > 0 {
				z01.PrintRune(vowels[vi])
				vi++
			} else {
				z01.PrintRune(r)
			}
		}
		if ai != len(args)-1 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
