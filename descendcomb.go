package piscine

import (
	"github.com/01-edu/z01"
)

func DescendComb() {
	for a := '9'; a >= '0'; a-- {
		for b := '9'; b >= '0'; b-- {
			for c := '9'; c >= '0'; c-- {
				for d := '9'; d >= '0'; d-- {
					if a < c || (b <= d && a == c) {
						continue
					}
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)

					if a == '0' && b == '1' && c == '0' && d == '0' {
						continue
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
