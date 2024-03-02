package main

import "fmt"

func main() {
	var n1, n2 float32
	fmt.Scan(&n1, &n2)

	if int(n1+n2) %  2 ==0 {
		fmt.Println("ЧЁТНОЕ")
	} else {
		fmt.Println("НЕЧЁТНОЕ")
	}
	
}
