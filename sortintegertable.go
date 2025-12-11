package piscine

// SortIntegerTable reorders a slice of int in ascending order without imports
func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if table[j] < table[minIdx] {
				minIdx = j
			}
		}
		table[i], table[minIdx] = table[minIdx], table[i]
	}
}
