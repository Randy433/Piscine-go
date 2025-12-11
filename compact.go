package piscine

func Compact(ptr *[]string) int {
	s := []string{}
	for i := 0; i < len(*ptr); i++ {
		if (*ptr)[i] != "" {
			s = append(s, (*ptr)[i])
		}
	}
	*ptr = []string{}
	for _, ch := range s {
		*ptr = append(*ptr, ch)
	}
	return len(*ptr)
}
