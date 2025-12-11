package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Convert from baseFrom to decimal (base 10)
	decimal := toDecimal(nbr, baseFrom)

	// Convert from decimal to baseTo
	return fromDecimal(decimal, baseTo)
}

// toDecimal converts a number string from any base to decimal
func toDecimal(nbr, base string) int {
	if nbr == "" || base == "" {
		return 0
	}

	result := 0
	baseLen := len(base)

	for _, char := range nbr {
		// Find the position of the character in the base
		pos := indexOf(base, char)
		if pos == -1 {
			return 0
		}

		result = result*baseLen + pos
	}

	return result
}

// fromDecimal converts a decimal number to any base
func fromDecimal(nbr int, base string) string {
	if nbr == 0 {
		return string(base[0])
	}

	baseLen := len(base)
	result := ""

	for nbr > 0 {
		remainder := nbr % baseLen
		result = string(base[remainder]) + result
		nbr = nbr / baseLen
	}

	return result
}

// indexOf finds the position of a rune in a string
func indexOf(s string, char rune) int {
	for i, c := range s {
		if c == char {
			return i
		}
	}
	return -1
}
