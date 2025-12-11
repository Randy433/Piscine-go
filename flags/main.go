package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 || args[0] == "--help" || args[0] == "-h" {
		printHelp()
		return
	}

	var insertStr string
	var orderFlag bool
	var mainStr string

	// Parse flags
	i := 0
	for i < len(args) {
		arg := args[i]

		if arg == "--insert" || arg == "-i" {
			if i+1 < len(args) {
				insertStr = args[i+1]
				i += 2
				continue
			}
		} else if arg == "--order" || arg == "-o" {
			orderFlag = true
			i++
			continue
		} else if len(arg) > 3 && arg[:3] == "-i=" {
			insertStr = arg[3:]
			i++
			continue
		} else if len(arg) > 9 && arg[:9] == "--insert=" {
			insertStr = arg[9:]
			i++
			continue
		} else {
			// This should be the main string
			mainStr = arg
			i++
			continue
		}
		i++
	}

	// Apply insertion if requested
	if insertStr != "" {
		mainStr += insertStr
	}

	// Apply ordering if requested
	if orderFlag {
		mainStr = sortString(mainStr)
	}

	// Print the result
	for _, r := range mainStr {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func printHelp() {
	// Print --insert and its description
	printLine("--insert")
	printLine("  -i")
	printLine("\t This flag inserts the string into the string passed as argument.")

	// Print --order and its description
	printLine("--order")
	printLine("  -o")
	printLine("\t This flag will behave like a boolean, if it is called it will order the argument.")
}

func printLine(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func sortString(s string) string {
	// Convert string to rune slice for sorting
	runes := []rune(s)

	// Simple bubble sort
	for i := 0; i < len(runes)-1; i++ {
		for j := 0; j < len(runes)-i-1; j++ {
			if runes[j] > runes[j+1] {
				runes[j], runes[j+1] = runes[j+1], runes[j]
			}
		}
	}

	return string(runes)
}
