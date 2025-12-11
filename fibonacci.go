package piscine

func Fibonacci(index int) int {
	if index < 0 {
		return -1
	} else if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	prev, curr := 0, 1
	for i := 2; i <= index; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
