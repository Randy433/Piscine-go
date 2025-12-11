package main

import (
	"fmt"
	"os"
)

func main() {
	// Check if we have enough arguments
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s -c <num> <file1> [file2 ...]\n", os.Args[0])
		os.Exit(1)
	}

	// Parse the -c option
	if os.Args[1] != "-c" {
		fmt.Fprintf(os.Stderr, "Error: expected -c option\n")
		os.Exit(1)
	}

	// Parse the number of bytes manually
	numBytes := parseNumber(os.Args[2])
	if numBytes < 0 {
		fmt.Fprintf(os.Stderr, "Error: invalid number of bytes\n")
		os.Exit(1)
	}

	// Get the list of files
	files := os.Args[3:]
	if len(files) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no files specified\n")
		os.Exit(1)
	}

	// Track if any errors occurred
	hasError := false
	multipleFiles := len(files) > 1
	printedFiles := 0

	// Process each file
	for i, filename := range files {
		// Try to open the file
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			hasError = true
			continue
		}

		// Print file header for multiple files (only after successful open)
		if multipleFiles {
			if i > 0 {
				fmt.Println() // Blank line between files
			}
			fmt.Printf("==> %s <==\n", filename)
		}

		// Read and print the last N bytes of the file
		err = tailFile(file, numBytes)
		file.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			hasError = true
		} else {
			printedFiles++
		}
	}

	// Exit with non-zero status if any errors occurred
	if hasError {
		os.Exit(1)
	}
}

// parseNumber converts a string to an integer without using strconv
func parseNumber(s string) int {
	if len(s) == 0 {
		return -1
	}

	result := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return -1 // Invalid character
		}
		result = result*10 + int(s[i]-'0')
	}

	return result
}

// tailFile reads and prints the last numBytes bytes from an open file
func tailFile(file *os.File, numBytes int) error {
	// Get file size
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}
	fileSize := fileInfo.Size()

	// Calculate how many bytes to read
	bytesToRead := numBytes
	if int64(bytesToRead) > fileSize {
		bytesToRead = int(fileSize)
	}

	// Seek to the position where we should start reading
	startPos := fileSize - int64(bytesToRead)
	_, err = file.Seek(startPos, 0)
	if err != nil {
		return err
	}

	// Read the last bytes
	buffer := make([]byte, bytesToRead)
	n, err := file.Read(buffer)
	if err != nil {
		return err
	}

	// Print the content
	fmt.Print(string(buffer[:n]))

	return nil
}
