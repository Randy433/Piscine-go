package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	n := len(a)

	// Bubble sort using custom comparison function
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Use the comparison function to determine order
			if f(a[j], a[j+1]) > 0 {
				// Swap if they're in wrong order
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

// Compare is a comparison function that compares strings by ASCII values (use "man ascii" in the terminal)
func Comparison(a, b string) int {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	// Compare each character by character
	for i := 0; i < minLen; i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}

	// If all characters match, shorter string comes first
	if len(a) < len(b) {
		return -1
	}
	if len(a) > len(b) {
		return 1
	}

	return 0
}
