package piscine

// Split takes a string s and a separator sep, and returns a slice of strings that results
// from splitting the string s by the separator sep.
func Split(s, sep string) []string {
	var result []string
	var currentString string

	// Traverse the string s character by character
	for i := 0; i < len(s); i++ {
		// Check if the current character is part of the separator
		if i+len(sep) <= len(s) && s[i:i+len(sep)] == sep {
			// If separator is found, add currentString to the result
			result = append(result, currentString)
			// Reset currentString for the next part
			currentString = ""
			// Skip ahead by the length of the separator
			i += len(sep) - 1
		} else {
			// Otherwise, add the character to the current word being formed
			currentString += string(s[i])
		}
	}

	// Append the last part of the string after the loop finishes
	result = append(result, currentString)

	return result
}
