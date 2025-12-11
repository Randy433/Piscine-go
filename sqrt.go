package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb <= 1 {
		return nb
	}
	for i := 1; i <= nb/i; i++ { // Avoids overflow, still checks i*i <= nb
		square := i * i
		if square == nb {
			return i
		}
		if square > nb {
			return i - 1 // Return largest int where i*i <= nb
		}
	}
	return 0
}
