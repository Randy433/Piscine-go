package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var order1 food
	var order2 food
	var order3 food

	order1.preptime = 15
	order2.preptime = 12
	order3.preptime = 10
	if order == "burger" {
		return order1.preptime
	} else if order == "nuggets" {
		return order2.preptime
	} else if order == "chips" {
		return order3.preptime
	} else {
		return 404
	}
}

// burger 15min
// chips 10min
// nuggets 12min
// if does not exit return 404
