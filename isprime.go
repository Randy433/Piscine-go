package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb <= 3 {
		return true // 2 and 3 are prime
	}
	if nb%2 == 0 || nb%3 == 0 {
		return false // Eliminate even numbers and multiples of 3
	}
	// Check only numbers of form 6k±1 up to √nb
	for i := 5; i*i <= nb; i += 6 {
		if nb%i == 0 || nb%(i+2) == 0 {
			return false
		}
	}
	return true
}

/**
simpler version if the code above
func IsPrime(nb int) bool {
    if nb <= 1 {
        return false
    }
    if nb == 2 {
        return true
    }
    if nb%2 == 0 {
        return false  // Quick check for even numbers
    }

    for i := 3; i*i <= nb; i += 2 {  // Skip even numbers
        if nb%i == 0 {
            return false
        }
    }
    return true
}
**/
