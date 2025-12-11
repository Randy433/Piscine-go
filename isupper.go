package piscine

func IsUpper(s string) bool {
	if s == "" {
		return false
	}
	for _, char := range s {
		if char < 'A' || char > 'Z' {
			return false
		}
	}
	return true
}
