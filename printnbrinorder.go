package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
	}
	var digits [10]int // declares an array of 10 integers (0-9) initialized to 0
	for ; n > 0; n /= 10 {
		digits[n%10]++
	}
	for i := 0; i < 10; i++ {
		for ; digits[i] > 0; digits[i]-- {
			z01.PrintRune(rune(i + '0')) // convert int to corresponding rune/ASCII character
		}
	}
}
