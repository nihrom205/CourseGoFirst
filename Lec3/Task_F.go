package main

import "fmt"

func main() {
	var (
		s1 string
		s2 string
		s3 string
		s4 string
		author string
	)

	fmt.Scan(&s1, &s2, &s3, &s4, &author)

	fmt.Printf("%s - %s\n", s4, author)
	fmt.Printf("%s - %s\n", s3, author)
	fmt.Printf("%s - %s\n", s2, author)
	fmt.Printf("%s - %s", s1, author)
}
