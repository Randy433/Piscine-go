package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	// Check for --upper flag
	upperCase := false
	startIndex := 0

	if args[0] == "--upper" {
		upperCase = true
		startIndex = 1
	}

	// Process each argument
	for i := startIndex; i < len(args); i++ {
		n := 0
		valid := true

		// Convert argument to integer
		for _, char := range args[i] {
			if char < '0' || char > '9' {
				valid = false
				break
			}
			n = n*10 + int(char-'0')
		}

		// Check if valid position in alphabet
		if valid && n >= 1 && n <= 26 {
			if upperCase {
				z01.PrintRune(rune('A' + n - 1))
			} else {
				z01.PrintRune(rune('a' + n - 1))
			}
		} else {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}
