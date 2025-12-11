package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arr := make([]int, max-min)
	for i := min; i < max; i++ {
		arr[i-min] = i
	}
	return arr
}
