package piscine

func Map(f func(int) bool, a []int) []bool {
	res := make([]bool, len(a))
	for i, v := range a {
		res[i] = f(v)
	}
	return res
}
