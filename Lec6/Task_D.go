package main

import (
	"fmt"
	"os"
)

func main() {
	var str string
	for {
		fmt.Fscan(os.Stdin, &str)

		fmt.Println(len(str))
		if str == "\n" {
			return
		}
		fmt.Println(str)
	}
}
