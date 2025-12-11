package piscine

func UltimateDivMod(a *int, b *int) {
	// storing the dereference
	div := *a / *b
	mod := *a % *b

	// pointing back the values
	*a = div
	*b = mod
}
