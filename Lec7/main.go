package main

import "fmt"

func main() {
	var price int
	fmt.Scan(&price)
	// switch

	switch price {
	case 100:
		fmt.Println("First case")
	case 110:
		fmt.Println("Second case")
	case 120:
		fmt.Println("Third case")
	case 130:
		fmt.Println("Another case")
	default:
		fmt.Println("Default case")
	}

	// case с множеством вариантов
	var ageGroup string = "dd" // возврастные группы: "a", "b", "c", "d", "e"
	switch ageGroup {
	case "a", "b", "c":
		fmt.Println("AgeGroup  10-40")
	case "d", "e":
		fmt.Println("AgeGroup  50-70")
	default:
		fmt.Println("yong")
	}


	// case с выражениями
	var age int
	fmt.Scan(&age)
	
	switch {
	case age <= 18:
		fmt.Println("too yong")
	case age > 18 && age <= 30:
		fmt.Println("Case second")
	case age > 30 && age <= 100:
		fmt.Println("too old")
	default:
		fmt.Println("who are you")
	}


	// case с проавливанием
	var num int
	fmt.Scan(&num)

	switch {
	case num < 100:
		fmt.Printf("%d is less then 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less then 200\n", num)
	}
}
