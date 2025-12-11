package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	// Check if base is valid
	if len(base) < 2 || hasDuplicates(base) || isIn(base, '-') || isIn(base, '+') {
		for _, v := range "NV" {
			z01.PrintRune(v)
		}
		return
	}
	// Handle 0 case
	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}
	isNeg := nbr < 0
	if isNeg {
		// Handle potential overflow for minimum int value
		if nbr == -9223372036854775808 {
			// Special handling for min int value
			handleMinValue(base)
			return
		}
		nbr = -nbr
		z01.PrintRune('-')
	}
	baseConverter(nbr, base)
}

func isIn(base string, r rune) bool {
	for _, v := range base {
		if v == r {
			return true
		}
	}
	return false
}

func hasDuplicates(base string) bool {
	// Check for duplicate characters in the base
	charMap := make(map[rune]bool)
	for _, char := range base {
		if charMap[char] {
			return true
		}
		charMap[char] = true
	}
	return false
}

func baseConverter(n int, base string) {
	lenbase := len(base)
	if n >= lenbase {
		baseConverter(n/lenbase, base)
	}
	z01.PrintRune(rune(base[n%lenbase]))
}

func handleMinValue(base string) {
	// Special handling for minimum int value to avoid overflow
	maxVal := 9223372036854775807
	lenbase := len(base)
	z01.PrintRune('-')
	// Convert the maximum value divided by base length to avoid overflow
	if maxVal/lenbase > 0 {
		baseConverter(maxVal/lenbase, base)
	}
	// Handle the remainder
	remainder := maxVal%lenbase + 1 // +1 because min value is max value + 1 in magnitude
	if remainder >= lenbase {
		// This needs adjustment for certain bases
		baseConverter(remainder/lenbase, base)
		remainder %= lenbase
	}
	z01.PrintRune(rune(base[remainder]))
}
