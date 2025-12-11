package piscine

func StringToIntSlice(str string) []int {
	var num []int
	for _, n := range str {
		num = append(num, int(n))
	}
	return num
}
