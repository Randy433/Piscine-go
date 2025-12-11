package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	v := os.Args[0]
	r := []rune(v)
	for i := 2; i < len(r); i++ {
		z01.PrintRune(r[i])
	}
	z01.PrintRune('\n')
}
