package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	}
	if s[0] == '+' {
		start = 1
	}

	rus := 0
	for i := start; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			rus = rus*10 + int(rune(s[i]-'0'))
		} else {
			return 0
		}
	}
	return rus * sign
}
