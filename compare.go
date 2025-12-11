package piscine

func Compare(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
