package main

import (
	"fmt"
	"math"
)

func main() {
	var count int
	var puls float64
	var max float64
	var min float64 = math.MaxFloat64 
	
	for {
		fmt.Scan(&puls)
		if puls == -1 {
			break
		}
		if puls >= 100 && puls <= 140 {
			count++
			if puls > max {
				max = puls
			}
			if puls < min {
				min = puls
			}
		}
	}
	fmt.Printf("%d\n%.1f %.1f", count, min, max)
}
