package piscine

// import "github.com/01-edu/z01"

func Rot14(s string) string {
	result := ""
	for _, c := range s {
		if c >= 'a' && c <= 'z' { // if c is between a and z
			if c+14 > 'z' {
				result = result + string(c+14-26) // when greater than z after changing position, take 26 positions from it
			} else {
				result = result + string(c+14)
			}
		} else if c >= 'A' && c <= 'Z' {
			if c+14 > 'Z' {
				result = result + string(c+14-26) // when greater than z after changing position, take 26 positions from it
			} else {
				result = result + string(c+14)
			}
		} else {
			result = result + string(c)
		}
	}
	return result
}
