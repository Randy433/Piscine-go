package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb <= 1 {
		return 1
	}

	result := 1
	for i := 2; i <= nb; i++ {
		result *= i
	}
	return result
}
