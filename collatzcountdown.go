package piscine

func CollatzCountdown(start int) int {
	counter := 0 // our countdown counter
	if start > 0 {
		// loop could have been i := start; i > 1; i-- or maybe use a range of "start"
		for start > 1 {
			if start%2 == 0 { // if n
				counter++
				start = start / 2
			} else {
				counter++
				start = (start * 3) + 1
			}
		}
	} else {
		return -1 // if start <= 0
	}
	return counter
}

// func main() {
// 	steps := CollatzCountdown(12)
// 	fmt.Println(steps)
// }

// link to the statement problem: https://en.wikipedia.org/wiki/Collatz_conjecture
