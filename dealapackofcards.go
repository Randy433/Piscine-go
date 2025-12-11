package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 0; i < 12; i += 3 {
		roundStart := i/3 + 1
		roundEnd := i + 3

		fmt.Printf("Player %v: ", roundStart)

		for i, round := range deck[i:roundEnd] {
			fmt.Printf("%v", round)
			if i != 2 {
				fmt.Printf(", ")
			}
		}
		fmt.Printf("\n")
	}
}
