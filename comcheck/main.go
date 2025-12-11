package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	count := 0

	for _, r := range arg {
		if r == "01" || r == "galaxy" || r == "galaxy 01" {
			count++
		}
	}
	if count > 0 {
		fmt.Println("Alert!!!")
	}
}
