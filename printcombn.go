package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	// Start with the smallest possible combination
	digits := make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = i
	}

	// Generate and print combinations
	for {
		// Print current combination
		for i := 0; i < n; i++ {
			z01.PrintRune(rune('0' + digits[i]))
		}

		// Find the rightmost digit that can be incremented
		pos := -1
		for i := n - 1; i >= 0; i-- {
			if digits[i] < 10-n+i {
				pos = i
				break
			}
		}

		// If no digit can be incremented, we're done
		if pos == -1 {
			break
		}

		// Print comma and space before next combination
		z01.PrintRune(',')
		z01.PrintRune(' ')

		// Increment the found digit and reset all digits to the right
		digits[pos]++
		for i := pos + 1; i < n; i++ {
			digits[i] = digits[i-1] + 1
		}
	}

	z01.PrintRune('\n')
}
