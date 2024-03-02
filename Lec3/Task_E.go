package main

import "fmt"

func main() {
	var (
		drink string
		meal string
		subMeal string
		time string
	)

	fmt.Scan(&drink, &meal, &subMeal, &time)

	// I wanna some 'tea', my favorite meal - 'chicken'. Give me also 'fries'. I will come soon... '13:00'
	fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'", drink, meal, subMeal, time)
}
