package main

import (
	"os"
)

func main() {
	// Check for exactly 3 arguments
	if len(os.Args) != 4 {
		return
	}

	// Parse the arguments
	value1, valid1 := parseNumber(os.Args[1])
	operator := os.Args[2]
	value2, valid2 := parseNumber(os.Args[3])

	// Check if values are valid
	if !valid1 || !valid2 {
		return
	}

	// Perform the operation
	var result int64
	var overflow bool

	switch operator {
	case "+":
		result, overflow = add(value1, value2)
		if overflow {
			return
		}
	case "-":
		result, overflow = subtract(value1, value2)
		if overflow {
			return
		}
	case "*":
		result, overflow = multiply(value1, value2)
		if overflow {
			return
		}
	case "/":
		if value2 == 0 {
			os.Stdout.WriteString("No division by 0\n")
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		result = value1 % value2
	default:
		return // Invalid operator
	}

	printNumber(result)
}

// parseNumber converts a string to int64, checking for overflow
func parseNumber(s string) (int64, bool) {
	if len(s) == 0 {
		return 0, false
	}

	negative := false
	startIdx := 0

	// Check for sign
	if s[0] == '-' {
		negative = true
		startIdx = 1
	} else if s[0] == '+' {
		startIdx = 1
	}

	// Check if there are digits after the sign
	if startIdx >= len(s) {
		return 0, false
	}

	var result int64 = 0
	maxInt64 := int64(^uint64(0) >> 1)
	minInt64 := -maxInt64 - 1

	for i := startIdx; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}

		digit := int64(s[i] - '0')

		// Check for overflow before adding the digit
		if negative {
			// For negative numbers, check against minInt64
			if result < (minInt64+digit)/10 {
				return 0, false
			}
			result = result*10 - digit
		} else {
			// For positive numbers, check against maxInt64
			if result > (maxInt64-digit)/10 {
				return 0, false
			}
			result = result*10 + digit
		}
	}

	return result, true
}

// printNumber prints an int64 to stdout
func printNumber(n int64) {
	if n == 0 {
		os.Stdout.WriteString("0\n")
		return
	}

	// Handle negative numbers
	if n < 0 {
		os.Stdout.WriteString("-")
		if n == int64(-9223372036854775808) {
			// Special case for most negative int64
			os.Stdout.WriteString("9223372036854775808\n")
			return
		}
		n = -n
	}

	// Convert number to string
	digits := []byte{}
	for n > 0 {
		digits = append(digits, byte('0'+n%10))
		n /= 10
	}

	// Reverse and print
	for i := len(digits) - 1; i >= 0; i-- {
		os.Stdout.Write([]byte{digits[i]})
	}
	os.Stdout.WriteString("\n")
}

// add performs addition with overflow detection
func add(a, b int64) (int64, bool) {
	maxInt64 := int64(^uint64(0) >> 1)
	minInt64 := -maxInt64 - 1

	result := a + b

	// Check for overflow
	if b > 0 && a > maxInt64-b {
		return 0, true
	}
	if b < 0 && a < minInt64-b {
		return 0, true
	}

	return result, false
}

// subtract performs subtraction with overflow detection
func subtract(a, b int64) (int64, bool) {
	maxInt64 := int64(^uint64(0) >> 1)
	minInt64 := -maxInt64 - 1

	result := a - b

	// Check for overflow
	if b < 0 && a > maxInt64+b {
		return 0, true
	}
	if b > 0 && a < minInt64+b {
		return 0, true
	}

	return result, false
}

// multiply performs multiplication with overflow detection
func multiply(a, b int64) (int64, bool) {
	minInt64 := int64(-9223372036854775808)

	if a == 0 || b == 0 {
		return 0, false
	}

	result := a * b

	// Check for overflow by dividing back
	if result/a != b {
		return 0, true
	}

	// Additional check for edge case: minInt64 * -1
	if a == minInt64 && b == -1 {
		return 0, true
	}
	if b == minInt64 && a == -1 {
		return 0, true
	}

	return result, false
}
