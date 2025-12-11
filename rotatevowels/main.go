package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	// If no arguments, print new line and return
	if len(args) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Join all arguments into a single string for vowel processing
	var allArgs string
	for _, arg := range args {
		allArgs += arg
	}

	// Check if there are any vowels
	hasVowels := false
	for _, r := range allArgs {
		if isVowel(r) {
			hasVowels = true
			break
		}
	}

	// If no vowels, print arguments with spaces
	if !hasVowels {
		for i, arg := range args {
			for _, r := range arg {
				z01.PrintRune(r)
			}
			if i < len(args)-1 {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
		return
	}

	// Find all vowel positions and collect vowels
	var vowelPositions []int
	var vowels []rune

	for i, r := range allArgs {
		if isVowel(r) {
			vowelPositions = append(vowelPositions, i)
			vowels = append(vowels, r)
		}
	}

	// Mirror the vowels
	for i := 0; i < len(vowels)/2; i++ {
		j := len(vowels) - 1 - i
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Reconstruct the string with mirrored vowels
	result := []rune(allArgs)
	for idx, pos := range vowelPositions {
		result[pos] = vowels[idx]
	}

	// Calculate where each argument ends in the combined string
	argLengths := make([]int, len(args))
	for i, arg := range args {
		argLengths[i] = len(arg)
	}

	// Print the result with spaces between arguments
	start := 0
	for i, length := range argLengths {
		// Print this argument
		for j := 0; j < length; j++ {
			z01.PrintRune(result[start+j])
		}
		// Print space if not the last argument
		if i < len(args)-1 {
			z01.PrintRune(' ')
		}
		start += length
	}
	z01.PrintRune('\n')
}

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}
