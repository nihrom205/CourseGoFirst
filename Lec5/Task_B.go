package main

import (
	"fmt"
)

func main() {
	var x, y, z float64
	fmt.Scan(&x, &y, &z)

	if x == 0 {
		if y == 0 {
			fmt.Println("корней нет")
		} else {
			fmt.Println("один корень")
		}
	} else {
		D := y*y - 4*x*z
		if D > 0 {
			fmt.Println("два корня")
		} else if D == 0 {
			fmt.Println("один корень")
		} else {
			fmt.Println("корней нет")
		}
	}
	
}
