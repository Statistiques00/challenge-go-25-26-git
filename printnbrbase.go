package student

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	if nbr < 0 {
		z01.PrintRune('-')
		// Handle MinInt edge case: use negative division/modulus
		printBaseNegative(nbr, base)
		return
	}
	printBase(nbr, base)
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	seen := make(map[rune]bool)
	for _, r := range base {
		if r == '+' || r == '-' || seen[r] {
			return false
		}
		seen[r] = true
	}
	return true
}

func printBase(n int, base string) {
	b := len(base)
	if n >= b {
		printBase(n/b, base)
	}
	z01.PrintRune(rune(base[n%b]))
}

func printBaseNegative(n int, base string) {
	b := len(base)
	if n <= -b {
		printBaseNegative(n/b, base)
	}
	mod := -(n % b)
	if mod < 0 {
		mod += b
	}
	z01.PrintRune(rune(base[mod]))
}
