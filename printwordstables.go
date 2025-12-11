package piscine

import (
	"github.com/01-edu/z01"
)

// prints each string in the slice on a separate line
func PrintWordsTables(a []string) {
	for _, word := range a {
		// Fprint each character on a new line using PrintRune from the z01 package
		for _, char := range word {
			z01.PrintRune(char)
		}
		// prints all characters of a word, print a newline
		z01.PrintRune('\n')
	}
}
