package piscine

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	sLen := len(s)
	findLen := len(toFind)
	if findLen > sLen {
		return -1
	}
	for i := 0; i <= sLen-findLen; i++ {
		if s[i:i+findLen] == toFind {
			return i
		}
	}
	return -1
}

// this is the inbuilf index function, but you have to import the "strings" module
//	func Index(s string, toFind string) int {
//		return strings.Index(s, toFind)
//	}
