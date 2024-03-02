package main

import (
	"fmt"
)

func main() {
	var one, two, three string
	fmt.Scan(&one, &two, &three)

	if one == "раз" && two == "два" && three == "три" {
		fmt.Println("ОК")
	} else if one == "один" && two == "два" && three == "три" {
		fmt.Println("ОК")
	} else if one == "1" && two == "2" && three == "3" {
		fmt.Println("ОК")
	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}
}
