package piscine

func Max(a []int) int {
	maxNum := 0
	for i := 0; i < len(a); i++ {
		if a[i] > maxNum {
			maxNum = a[i]
		}
	}
	return maxNum
}
