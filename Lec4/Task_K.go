package main

import (
	"fmt"
	"strings"
)

func main() {
	bot := "чат"
	var text string
	fmt.Scan(&text)

	if strings.Contains(text, bot) {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}
}
