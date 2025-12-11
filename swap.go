package piscine

/*
  - func Swap(a *int, b *int) {
    *a, *b = *b, *a

} *
*/
func Swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
}
