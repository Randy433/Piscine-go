package piscine

func IsLower(s string) bool {
	if s == "" {
		return false
	}
	for _, char := range s {
		if char < 'a' || char > 'z' {
			return false
		}
	}
	return true
}
