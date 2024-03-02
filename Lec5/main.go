package main

import (
	"fmt"
	"strings"
)

func main() {
	var value int
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("The number", value, "is even")
	} else {
		fmt.Println("The number", value, "is odd")
	}

	var color string
	fmt.Scan(&color)
	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknnow color")
	}


	// инициализация в блоке условного оператора
	// блок присваивания - только :=
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")
	}

	// Идеоматичность
	if height := 100; height > 100 {
		fmt.Println("height > 100")
		return
	}
	fmt.Println("height <= 100")
}
