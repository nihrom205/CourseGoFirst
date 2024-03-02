package main

import (
	"fmt"
	"strings"
)

func main() {
	outer:	// лэйблы
		for i := 0; i < 2; i++ {
			for j := 1; j <= 3; j++ {
				fmt.Printf("i:%d and j:%d and sum i+j=%d\n", i,j,i+j)
				if i == j {
					break outer
				}
			}
		}

		var password string
		for {
			fmt.Println("Insert password")
			fmt.Scan(&password)
			if strings.Contains(password, "1234") {
				fmt.Println("Weak password. Try again")
			} else {
				fmt.Println("Password Accepted")
				break
			}
		}
}