package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summ := make(map[string]int)
	items := ""
	for _, ch := range str {
		if ch == ' ' {
			summ[items]++
			items = ""
		} else {
			items += string(ch)
		}
	}
	// Add the last item
	summ[items]++
	return summ
}
