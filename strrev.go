package piscine

func StrRev(s string) string {
	result := ""
	for _, char := range s {
		result = string(char) + result
	}
	return result
}
