package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arr := os.Args[1:]
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	for _, w := range arr {
		for _, c := range w {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
