package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Scan(&x1, &y1, &x2, &y2)

	deltaX := math.Abs(float64(x2) - float64(x1))
	deltaY := math.Abs(float64(y2) - float64(y1))

	if (deltaX == 1 && deltaY == 2) || (deltaX == 2 && deltaY == 1) {
		fmt.Println("ДА")
	} else {
		fmt.Println("НЕТ")
	}
}
