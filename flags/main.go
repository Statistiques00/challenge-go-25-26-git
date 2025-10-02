package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func printHelp() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("\t This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}
	insertStr := ""
	order := false
	argStr := ""
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			printHelp()
			return
		} else if arg == "--order" || arg == "-o" {
			order = true
		} else if len(arg) > 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
		} else if arg != "--insert" && arg != "-i" {
			argStr = arg
		}
	}
	if insertStr != "" {
		argStr += insertStr
	}
	if order {
		runes := []rune(argStr)
		for i := 0; i < len(runes)-1; i++ {
			for j := 0; j < len(runes)-i-1; j++ {
				if runes[j] > runes[j+1] {
					runes[j], runes[j+1] = runes[j+1], runes[j]
				}
			}
		}
		argStr = string(runes)
	}
	for _, r := range argStr {
		z01.PrintRune(r)
	}
	if argStr != "" {
		z01.PrintRune('\n')
	}
}
