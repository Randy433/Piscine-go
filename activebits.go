package piscine

func ActiveBits(n int) int {
	bit := 0
	for n != 0 {
		bit = bit + n&1
		n = n >> 1
	}
	return bit
}
