package student

import "github.com/01-edu/z01"

func PrintNbr(nb int) {
	if nb == 0 {
		z01.PrintRune('0')
		return
	}
	if nb < 0 {
		z01.PrintRune('-')
		if nb == -9223372036854775808 {
			for _, r := range "-9223372036854775808"[1:] {
				z01.PrintRune(r)
			}
			return
		}
		if nb == -2147483648 {
			for _, r := range "-2147483648"[1:] {
				z01.PrintRune(r)
			}
			return
		}
		nb = -nb
	}
	digits := []int{}
	for nb > 0 {
		digits = append(digits, nb%10)
		nb /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(rune(digits[i] + '0'))
	}
}
