package piscine

func IsPrintable(s string) bool {
	for _, r := range s {
		if r < 32 || r > 126 {
			return false
		}
	}
	return true
}

// Printable ASCII characters are those in the range from 32 to 126 inclusive.
