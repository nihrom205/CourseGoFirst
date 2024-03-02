package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Scan(&n1, &n2)

	fmt.Printf("Периметр: %d\nПлощадь: %d", n1+n1 + n2+n2, n1*n2)
}
