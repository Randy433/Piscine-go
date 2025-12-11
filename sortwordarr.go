package piscine

func SortWordArr(a []string) {
	n := len(a)

	// Bubble sort algorithm
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			// Compare adjacent strings
			if compareStrings(a[j], a[j+1]) > 0 {
				// Swap if they're in wrong order
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

// compareStrings compares two strings by ASCII values (use man ASCII in terminal)
// Returns: -1 if s1 < s2, 0 if s1 == s2, 1 if s1 > s2
func compareStrings(s1, s2 string) int {
	minLen := len(s1)
	if len(s2) < minLen {
		minLen = len(s2)
	}

	// Compare each character by character
	for i := 0; i < minLen; i++ {
		if s1[i] < s2[i] {
			return -1
		}
		if s1[i] > s2[i] {
			return 1
		}
	}

	// If and when all characters match, shorter string comes first
	if len(s1) < len(s2) {
		return -1
	}
	if len(s1) > len(s2) {
		return 1
	}

	return 0
}
