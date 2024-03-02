package main

import "fmt"

func main()  {
	var num int
	for {
		fmt.Scan(&num)
		if num == 0 {
			return
		}
		fmt.Println(num)
	}
}