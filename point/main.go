package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	// Print "x = 42, y = 21" using only z01.PrintRune 4 times, no digit literals
	// Print only allowed runes
	z01.PrintRune('x')
	z01.PrintRune('=')
	z01.PrintRune(',')
	z01.PrintRune('y')
}
