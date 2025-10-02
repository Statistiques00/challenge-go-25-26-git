package student

import "github.com/01-edu/z01"

const size = 8

func EightQueens() {
	var queens [size]int
	var cols, d1, d2 [2 * size]bool
	solveQueens(0, &queens, &cols, &d1, &d2)
}

func solveQueens(row int, queens *[size]int, cols, d1, d2 *[2 * size]bool) {
	if row == size {
		for i := 0; i < size; i++ {
			z01.PrintRune(rune(queens[i] + '0'))
		}
		z01.PrintRune('\n')
		return
	}
	for col := 0; col < size; col++ {
		if cols[col] || d1[row-col+size] || d2[row+col] {
			continue
		}
		queens[row] = col + 1
		cols[col], d1[row-col+size], d2[row+col] = true, true, true
		solveQueens(row+1, queens, cols, d1, d2)
		cols[col], d1[row-col+size], d2[row+col] = false, false, false
	}
}
