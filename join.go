package piscine

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return ""
	}

	result := ""
	for i, s := range strs {
		result += s
		if i < len(strs)-1 {
			result += sep
		}
	}
	return result
}
