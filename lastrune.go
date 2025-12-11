package piscine

func LastRune(s string) rune {
	cc := []rune(s)
	return cc[len(cc)-1]
}
