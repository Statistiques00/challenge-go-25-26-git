package student

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	digits := []rune("0123456789")
	for a := 0; a <= 98; a++ {
		for b := a + 1; b <= 99; b++ {
			z01.PrintRune(digits[a/10])
			z01.PrintRune(digits[a%10])
			z01.PrintRune(' ')
			z01.PrintRune(digits[b/10])
			z01.PrintRune(digits[b%10])
			if a != 98 || b != 99 {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
	z01.PrintRune('\n')
}
