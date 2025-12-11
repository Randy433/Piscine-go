package piscine

func SplitWhiteSpaces(s string) []string {
	r := []string{}
	index := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if s[index] != ' ' {
				r = append(r, s[index:i])
			}
			index = i + 1
		} else if i == len(s)-1 {
			r = append(r, s[index:])
		}
	}
	return r
}
