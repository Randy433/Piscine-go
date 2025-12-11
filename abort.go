package piscine

// import (
// 	"slices"
// )

func sortDigits2(digits []int) {
	length := len(digits)

	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if digits[j] > digits[j+1] {
				digits[j], digits[j+1] = digits[j+1], digits[j]
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	sortDigits2(arr)
	// slices.Sort(arr)
	abc := (len(arr) + 1) / 2
	return arr[abc-1]
}

// simple solution for this: return (a+b+c+d+e)/2
