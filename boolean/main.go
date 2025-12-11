package main

import (
	"os"

	"github.com/01-edu/z01"
)

// declare a type boolean
type boolean int

// use  constants for boolean values and messages
const (
	yes     boolean = 1
	no      boolean = 0
	EvenMsg string  = "I have an even number of arguments"
	OddMsg  string  = "I have an odd number of arguments"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func even(nbr int) int {
	if nbr%2 == 0 {
		return 1
	}
	return 0
}

func isEven(nbr int) boolean {
	if even(nbr) == 1 {
		return yes
	} else {
		return no
	}
}

func main() {
	lengthOfArg := len(os.Args[1:])
	if isEven(lengthOfArg) == 1 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
