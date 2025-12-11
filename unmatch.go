package piscine

func Unmatch(a []int) int {
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		counter := 0

		for j := 0; j < aLen; j++ {
			if i != j && a[i] == a[j] {
				counter++
			}
		}

		if counter == 0 || counter%2 == 0 {
			return a[i]
		}
	}
	return -1
}

// func main() {
// 	a := []int{1, 2, 3, 1, 2, 3, 4}
// 	unmatch := Unmatch(a)
// 	fmt.Println(unmatch)
// }
